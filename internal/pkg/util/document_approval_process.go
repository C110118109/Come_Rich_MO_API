package util

//這份檔案定義簽核流程規則
//完整規劃目錄放置於 https://hackmd.io/@bovcu13/SyEFBZFpj/edit
// Process design by Vinchee, jjps30516@gmail.com

func ProcessRule(isInit bool, routerName string, role string, isApprocal bool) (topic string, nextRole string) {

	// routerName請填寫流程名稱
	// return 下一個任務的對象跟名稱 (isApprocal代表通過/ !isApprocal代表不通過

	//訂單部分 PISP v1.1 (20230302 pass)
	if routerName == "order" {

		if isInit { //起單動作
			return "業務複審", "Sales_EMP"

		} else { //一般審核動作

			//複審動作
			if role == "Sales_EMP" && isApprocal {
				return "業務主管審核", "Sales_SIR"
			} else if role == "Sales_EMP" && !isApprocal {
				return "退回(業務複審未通過)", "Sales_EMP"
			}

			//進入業務主管審核
			if role == "Sales_SIR" && isApprocal {
				return "財務經辦審核", "Finance_ATTN"
			} else if role == "Sales_SIR" && !isApprocal {
				return "退回(業務主管審核未通過)", "Sales_EMP"
			}

			//進入財務經辦審核
			if role == "Finance_ATTN" && isApprocal {
				return "財務主管審核", "Finance_SIR"
			} else if role == "Finance_ATTN" && !isApprocal {
				return "退回至業務端(財務經辦審核未通過)", "Sales_EMP"
			}

			//進入財務主管審核
			if role == "Finance_SIR" && isApprocal {
				return "總經理審核", "GM"
			} else if role == "Finance_SIR" && !isApprocal {
				return "退回至業務端(財務主管審核未通過)", "Sales_EMP"
			}

			//進入總經理審核
			if role == "GM" && isApprocal {
				return "", "Done"
			} else if role == "GM" && !isApprocal {
				return "退回至業務端(總經理審核未通過)", "Sales_EMP"
			}
		}
	}

	//進貨單 PO
	if routerName == "purchase-order" {

		if isInit { //啟單動作
			return "業務複審", "Sales_EMP"

		} else { //一般審核動作

			//複審動作
			if role == "Sales_EMP" && isApprocal {
				return "業務主管審核", "Sales_SIR"
			} else if role == "Sales_EMP" && !isApprocal {
				return "退回(業務複審未通過)", "Sales_EMP"
			}

			//進入業務主管審核
			if role == "Sales_SIR" && isApprocal {
				return "", "Done"
			} else if role == "Sales_SIR" && !isApprocal {
				return "退回(業務主管審核未通過)", "Sales_EMP"
			}
		}
	}

	//請款單
	if routerName == "invoice" {

		if isInit { //啟單動作
			return "船務審核", "Shipping_EMP"

		} else { //一般審核動作

			//進入船務審核
			if role == "Shipping_EMP" && isApprocal {
				return "業務主管審核", "Sales_SIR"
			} else if role == "Shipping_EMP" && !isApprocal {
				return "退回至業務端(船務審核未通過)", "Sales_EMP"
			}

			//進入業務主管審核
			if role == "Sales_SIR" && isApprocal {
				return "財務經辦審核", "Finance_ATTN"
			} else if role == "Sales_SIR" && !isApprocal {
				return "退回(業務主管審核未通過)", "Sales_EMP"
			}

			//進入財務經辦審核
			if role == "Finance_ATTN" && isApprocal {
				return "財務主管審核", "Finance_SIR"
			} else if role == "Finance_ATTN" && !isApprocal {
				return "退回至業務端(財務經辦審核未通過)", "Sales_EMP"
			}

			//進入財務主管審核
			if role == "Finance_SIR" && isApprocal {
				return "總經理審核", "GM"
			} else if role == "Finance_SIR" && !isApprocal {
				return "退回至業務端(財務主管審核未通過)", "Sales_EMP"
			}

			//進入總經理審核
			if role == "GM" && isApprocal {
				return "財務付款", "Finance_EMP"
			} else if role == "GM" && !isApprocal {
				return "退回至業務端(總經理審核未通過)", "Sales_EMP"
			}

			//進入財務付款
			if role == "Finance_EMP" && isApprocal {
				return "", "Done"
			}
		}
	}

	//提單(分成兩個part  財務共通 船務個別)
	if routerName == "commercial-invoice" {

		if isInit { //啟單動作
			return "設定提單狀態(財務)", "Finance_"

		} else { //一般審核動作

			//設定提單狀態(由前端控制哪位角色可以完成送審)
			if role == "Finance_" && isApprocal {
				return "單據進行中", "Done"
			} else if role == "Finance_" && !isApprocal {
				return "退回(財務退回)", "Sales_EMP"
			}
		}
	}

	//提單2  For 個別船務
	if routerName == "commercial-invoice-individual" {

		if isInit { //啟單動作
			return "設定提單狀態(船務)", "Shipping_"

		} else { //一般審核動作

			//設定提單狀態(由前端控制哪位角色可以完成送審)(船務會另外依照使用者各自起單)
			if role == "Shipping_" && isApprocal {
				return "單據進行中", "Done"
			} else if role == "Shipping_" && !isApprocal {
				return "退回(船務退回)", "Sales_EMP"
			}
		}
	}

	//月差異表
	if routerName == "analysis-variances" {

		if isInit { //啟單動作
			return "業務主管審核", "Sales_SIR"

		} else { //一般審核動作

			//進入業務主管審核
			if role == "Sales_SIR" && isApprocal {
				return "總經理審核", "GM"
			} else if role == "Sales_SIR" && !isApprocal {
				return "退回(業務主管審核未通過)", "Sales_EMP"
			}

			//進入總經理審核
			if role == "GM" && isApprocal {
				return "", "Done"
			} else if role == "GM" && !isApprocal {
				return "退回至業務端(總經理審核未通過)", "Sales_EMP"
			}
		}
	}

	//結算表
	if routerName == "settlement-sheet" {

		if isInit { //啟單動作
			return "船務登記費用", "Shipping_EMP"

		} else { //一般審核動作

			//船務登記費用
			if role == "Shipping_EMP" && isApprocal {
				return "財務登記費用", "Finance_EMP"
			} else if role == "Shipping_EMP" && !isApprocal {
				return "退回(船務取消登記費用)", "Sales_EMP"
			}

			//財務登記費用
			if role == "Finance_EMP" && isApprocal {
				return "業務主管審核", "Sales_SIR"
			} else if role == "Finance_EMP" && !isApprocal {
				return "退回(財務取消登記費用)", "Sales_EMP"
			}

			//進入業務主管審核
			if role == "Sales_SIR" && isApprocal {
				return "財務主管審核", "Finance_SIR"
			} else if role == "Sales_SIR" && !isApprocal {
				return "退回(業務主管審核未通過)", "Sales_EMP"
			}

			//進入財務主管審核
			if role == "Finance_SIR" && isApprocal {
				return "總經理審核", "GM"
			} else if role == "Finance_SIR" && !isApprocal {
				return "退回至業務端(財務主管審核未通過)", "Sales_EMP"
			}

			//進入總經理審核
			if role == "GM" && isApprocal {
				return "", "Done"
			} else if role == "GM" && !isApprocal {
				return "退回至業務端(總經理審核未通過)", "Sales_EMP"
			}
		}
	}

	//收到客戶水單 (暫無使用)
	if routerName == "bank-slip-received" {

		if isInit { //啟單動作
			return "財務審核", "Finance_EMP"

		} else { //一般審核動作

			//進入財務審核
			if role == "Finance_EMP" && isApprocal {
				return "", "Done"
			} else if role == "Finance_EMP" && !isApprocal {
				return "退回至業務端(財務審核未通過)", "Sales_EMP"
			}
		}
	}

	return "", ""
}
