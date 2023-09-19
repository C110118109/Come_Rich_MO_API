package goods_detail

type Table struct {
	// 單據號碼
	SalesNo string `gorm:"primaryKey;column:SD_NO;type:TEXT" json:"sales_no,omitempty"`
	// 識別序號
	ItemSeq int `gorm:"column:SD_SEQ;type:int" json:"item_seq,omitempty"`
	// 商品名稱
	Name string `gorm:"column:SD_NAME;type:TEXT" json:"name,omitempty"`
	// 單位
	Unit string `gorm:"column:SD_UNIT;type:TEXT" json:"unit,omitempty"`
	// 價格
	Price float32 `gorm:"column:SD_PRICE;type:float" json:"price,omitempty"`
	// 數量
	Quantity float32 `gorm:"column:SD_QTY;type:float" json:"quantity,omitempty"`
	// 說明
	Remark string `gorm:"column:SD_BY1;type:TEXT" json:"remark,omitempty"`
}

type Base struct {
	// 顯示序號(客製)
	ItemNum int `json:"item_num,omitempty"`
	// 商品名稱
	Name string `json:"name"`
	// 數量
	Quantity float32 `json:"quantity"`
	// 單位
	Unit string `json:"unit"`
	// 單價
	Price float32 `json:"price"`
	// 小計(客製)
	Subtotal float64 `json:"subtotal,omitempty"`
	// 說明
	Remark string `json:"remark"`

	// 識別號碼(客製)
	Identifier string `json:"identifier,omitempty"`
	// 製令圖片編號(客製)
	FileId string `json:"file_id"`
	// 製令圖片路徑(客製)
	ImgPath string `json:"img_path"`
}

// TableName sets the insert table name for this struct type
func (c *Table) TableName() string {
	return "SSLPDT"
}
