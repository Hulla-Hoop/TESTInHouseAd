package goods

import "encoding/json"

func (s *service) CreateGoods(reqId string, name string) ([]byte, error) {
	category, err := s.db.CreateGoods(reqId, name)
	if err != nil {
		s.logger.L.WithField("SERVICE.CreateCategory", reqId).Error(err)
		return nil, err
	}
	s.logger.L.WithField("SERVICE.CreateCategory", reqId).Debug("Query row in service layer - ", category)
	SL, err := json.Marshal(category)
	if err != nil {
		s.logger.L.WithField("SERVICE.CreateCategory", reqId).Error(err)
		return nil, err
	}
	return SL, nil
}
