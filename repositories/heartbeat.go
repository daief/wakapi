package repositories

import (
	conf "github.com/muety/wakapi/config"
	"github.com/muety/wakapi/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type HeartbeatRepository struct {
	db     *gorm.DB
	config *conf.Config
}

func NewHeartbeatRepository(db *gorm.DB) *HeartbeatRepository {
	return &HeartbeatRepository{config: conf.Get(), db: db}
}

// Use with caution!!
func (r *HeartbeatRepository) GetAll() ([]*models.Heartbeat, error) {
	var heartbeats []*models.Heartbeat
	if err := r.db.Find(&heartbeats).Error; err != nil {
		return nil, err
	}
	return heartbeats, nil
}

func (r *HeartbeatRepository) InsertBatch(heartbeats []*models.Heartbeat) error {
	if err := r.db.
		Clauses(clause.OnConflict{
			DoNothing: true,
		}).
		Create(&heartbeats).Error; err != nil {
		return err
	}
	return nil
}

func (r *HeartbeatRepository) GetLatestByUser(user *models.User) (*models.Heartbeat, error) {
	var heartbeat models.Heartbeat
	if err := r.db.
		Model(&models.Heartbeat{}).
		Where(&models.Heartbeat{UserID: user.ID}).
		Order("time desc").
		First(&heartbeat).Error; err != nil {
		return nil, err
	}
	return &heartbeat, nil
}

func (r *HeartbeatRepository) GetLatestByOriginAndUser(origin string, user *models.User) (*models.Heartbeat, error) {
	var heartbeat models.Heartbeat
	if err := r.db.
		Model(&models.Heartbeat{}).
		Where(&models.Heartbeat{
			UserID: user.ID,
			Origin: origin,
		}).
		Order("time desc").
		First(&heartbeat).Error; err != nil {
		return nil, err
	}
	return &heartbeat, nil
}

func (r *HeartbeatRepository) GetAllWithin(from, to time.Time, user *models.User) ([]*models.Heartbeat, error) {
	// https://stackoverflow.com/a/20765152/3112139
	var heartbeats []*models.Heartbeat
	if err := r.db.
		Where(&models.Heartbeat{UserID: user.ID}).
		Where("time >= ?", from.Local()).
		Where("time < ?", to.Local()).
		Order("time asc").
		Find(&heartbeats).Error; err != nil {
		return nil, err
	}
	return heartbeats, nil
}

func (r *HeartbeatRepository) GetAllWithinByFilters(from, to time.Time, user *models.User, filterMap map[string][]string) ([]*models.Heartbeat, error) {
	// https://stackoverflow.com/a/20765152/3112139
	var heartbeats []*models.Heartbeat

	q := r.db.
		Where(&models.Heartbeat{UserID: user.ID}).
		Where("time >= ?", from.Local()).
		Where("time < ?", to.Local()).
		Order("time asc")

	for col, vals := range filterMap {
		q = q.Where(col+" in ?", vals)
	}

	if err := q.Find(&heartbeats).Error; err != nil {
		return nil, err
	}
	return heartbeats, nil
}

func (r *HeartbeatRepository) GetFirstByUsers() ([]*models.TimeByUser, error) {
	var result []*models.TimeByUser
	r.db.Model(&models.User{}).
		Select("users.id as user, min(time) as time").
		Joins("left join heartbeats on users.id = heartbeats.user_id").
		Group("user").
		Scan(&result)
	return result, nil
}

func (r *HeartbeatRepository) GetLastByUsers() ([]*models.TimeByUser, error) {
	var result []*models.TimeByUser
	r.db.Model(&models.User{}).
		Select("users.id as user, max(time) as time").
		Joins("left join heartbeats on users.id = heartbeats.user_id").
		Group("user").
		Scan(&result)
	return result, nil
}

func (r *HeartbeatRepository) Count(approximate bool) (count int64, err error) {
	if r.config.Db.IsMySQL() && approximate {
		err = r.db.Table("information_schema.tables").
			Select("table_rows").
			Where("table_schema = ?", r.config.Db.Name).
			Where("table_name = 'heartbeats'").
			Scan(&count).Error
	}

	if count == 0 {
		err = r.db.
			Model(&models.Heartbeat{}).
			Count(&count).Error
	}
	return count, nil
}

func (r *HeartbeatRepository) CountByUser(user *models.User) (int64, error) {
	var count int64
	if err := r.db.
		Model(&models.Heartbeat{}).
		Where(&models.Heartbeat{UserID: user.ID}).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *HeartbeatRepository) CountByUsers(users []*models.User) ([]*models.CountByUser, error) {
	var counts []*models.CountByUser

	userIds := make([]string, len(users))
	for i, u := range users {
		userIds[i] = u.ID
	}

	if len(userIds) == 0 {
		return counts, nil
	}

	if err := r.db.
		Model(&models.Heartbeat{}).
		Select("user_id as user, count(id) as count").
		Where("user_id in ?", userIds).
		Group("user").
		Find(&counts).Error; err != nil {
		return counts, err
	}

	return counts, nil
}

func (r HeartbeatRepository) GetEntitySetByUser(entityType uint8, user *models.User) ([]string, error) {
	var results []string
	if err := r.db.
		Model(&models.Heartbeat{}).
		Distinct(models.GetEntityColumn(entityType)).
		Where(&models.Heartbeat{UserID: user.ID}).
		Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (r *HeartbeatRepository) DeleteBefore(t time.Time) error {
	if err := r.db.
		Where("time <= ?", t.Local()).
		Delete(models.Heartbeat{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *HeartbeatRepository) DeleteByUser(user *models.User) error {
	if err := r.db.
		Where("user_id = ?", user.ID).
		Delete(models.Heartbeat{}).Error; err != nil {
		return err
	}
	return nil
}
