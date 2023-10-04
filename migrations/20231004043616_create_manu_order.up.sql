-- 製造指令單
CREATE TABLE manu_order
(
  manu_order_id    TEXT      NULL     ,
  goods_name       TEXT      NULL    ,
  identifier       TEXT      NULL     ,
  total_quantity   TEXT      NULL    ,
  completion_date  DATETIME  NULL    ,
  date_of_issuance DATETIME  NULL     ,
  created_at       TIMESTAMP NULL     ,
  updated_at       TIMESTAMP NULL     ,
  is_deleted	      BOOLEAN   NULL   ,
  PRIMARY KEY (manu_order_id) 
) ;