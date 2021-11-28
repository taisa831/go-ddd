package repository

import (
	"github.com/taisa831/go-ddd/domain/model"
	"gorm.io/gorm"
)

func (r *dbRepository) SaveCircle(c *model.Circle) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		record := r.convertToCircleRecord(c)
		err := tx.Model(circle{}).Where("id = ?", record.ID).Save(&record).Error
		if err != nil {
			err := tx.Model(circle{}).Save(&record).Error
			return err
		}

		memberRecords := r.convertToCircleMemberRecord(c.Members)
		for _, v := range memberRecords {
			err := tx.Model(circleMember{}).Where("user_id = ?", v.UserID).Save(&v).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *dbRepository) FindCircleByID(circleID *model.CircleID) (*model.Circle, error) {
	record := circle{}
	if err := r.db.Model(circle{}).Where("id = ?", circleID.Value).First(&record).Error; err != nil {
		return nil, err
	}
	return r.convertToCircleModel(record), nil
}

func (r *dbRepository) FindCircleByName(circleName *model.CircleName) (*model.Circle, error) {
	record := circle{}
	if err := r.db.Model(circle{}).Where("name = ?", circleName.Value).First(&record).Error; err != nil {
		return nil, err
	}
	return r.convertToCircleModel(record), nil
}

func (r *dbRepository) convertToCircleModel(c circle) *model.Circle {
	return &model.Circle{
		ID:    model.CircleID{*c.ID},
		Name:  model.CircleName{*c.Name},
		Owner: *c.Owner,
	}
}

func (r *dbRepository) convertToCircleRecord(c *model.Circle) *circle {
	return &circle{
		ID:    &c.ID.Value,
		Name:  &c.Name.Value,
		Owner: &c.Owner,
	}
}

func (r dbRepository) convertToCircleMemberRecord(members []*model.CircleMember) []*circleMember {
	circleMembers := []*circleMember{}
	for _, _m := range members {
		m := _m
		memberRecord := &circleMember{
			ID:       &m.ID,
			CircleID: &m.CircleID.Value,
			UserID:   &m.UserID,
		}
		circleMembers = append(circleMembers, memberRecord)
	}
	return circleMembers
}
