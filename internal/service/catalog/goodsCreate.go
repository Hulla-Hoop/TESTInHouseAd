package catalog

import "encoding/json"

func (s *catalog) CreateGoods(reqId string, name string) ([]byte, error) {
	category, err := s.db.CreateGoods(reqId, name)
	if err != nil {
		s.logger.L.WithField("SERVICE.CreateGoods", reqId).Error(err)
		return nil, err
	}
	s.logger.L.WithField("SERVICE.CreateGoods", reqId).Debug("Query row in catalog layer - ", category)
	SL, err := json.Marshal(category)
	if err != nil {
		s.logger.L.WithField("SERVICE.CreateGoods", reqId).Error(err)
		return nil, err
	}
	return SL, nil
}
