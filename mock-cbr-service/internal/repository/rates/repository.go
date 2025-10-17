package ratesrepo

import (
	"bytes"
	"fmt"
	"time"
)

type Repository struct{}

func New() *Repository {
	return &Repository{}
}

// GetDailyXML возвращает заглушку XML аналогичную ответу ЦБ для заданной даты
func (r *Repository) GetDailyXML(date time.Time) ([]byte, error) {
	xml := fmt.Sprintf(`<?xml version="1.0" encoding="windows-1251"?>
<ValCurs Date="%s" name="Foreign Currency Market">
  <Valute ID="R01235">
    <NumCode>840</NumCode>
    <CharCode>USD</CharCode>
    <Nominal>1</Nominal>
    <Name>Доллар США</Name>
    <Value>31,3500</Value>
  </Valute>
</ValCurs>`, date.Format("02.01.2006"))
	return bytes.NewBufferString(xml).Bytes(), nil
}
