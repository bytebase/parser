package postgresql

import (
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

type PostgreSQLLexerBase struct {
	*antlr.BaseLexer

	stack              StringStack
	reservedKeywordMap map[string]bool
}

func (receiver *PostgreSQLLexerBase) pushTag() {
	receiver.stack.Push(receiver.GetText())
}

func (receiver *PostgreSQLLexerBase) isTag() bool {
	if receiver.stack.IsEmpty() {
		return false
	}
	return receiver.GetText() == receiver.stack.PeekOrEmpty()
}

func (receiver *PostgreSQLLexerBase) popTag() {
	_, _ = receiver.stack.Pop()
}

func (receiver *PostgreSQLLexerBase) checkLA(c int) bool {
	return receiver.GetInputStream().LA(1) != c
}

func (receiver *PostgreSQLLexerBase) charIsLetter() bool {
	c := receiver.GetInputStream().LA(-1)
	return unicode.IsLetter(rune(c))
}

func (receiver *PostgreSQLLexerBase) HandleNumericFail() {
	index := receiver.GetInputStream().Index() - 2
	receiver.GetInputStream().Seek(index)
	receiver.SetType(PostgreSQLLexerIntegral)
}

func (receiver *PostgreSQLLexerBase) HandleLessLessGreaterGreater() {
	if receiver.GetText() == "<<" {
		receiver.SetType(PostgreSQLLexerLESS_LESS)
	}
	if receiver.GetText() == ">>" {
		receiver.SetType(PostgreSQLLexerGREATER_GREATER)
	}
}

func (receiver *PostgreSQLLexerBase) UnterminatedBlockCommentDebugAssert() {
	//Debug.Assert(InputStream.LA(1) == -1 /*EOF*/);
}

func (receiver *PostgreSQLLexerBase) CheckIfUtf32Letter() bool {
	codePoint := receiver.GetInputStream().LA(-2)<<8 + receiver.GetInputStream().LA(-1)
	var c []rune
	if codePoint < 0x10000 {
		c = []rune{rune(codePoint)}
	} else {
		codePoint -= 0x10000
		c = []rune{
			(rune)(codePoint/0x400 + 0xd800),
			(rune)(codePoint%0x400 + 0xdc00),
		}
	}
	return unicode.IsLetter(c[0])
}

func (receiver *PostgreSQLLexerBase) IsIdentifier(tokenType int) bool {
	if tokenType == antlr.TokenEOF {
		return false
	}

	switch tokenType {
	case PostgreSQLLexerIdentifier,
		PostgreSQLLexerQuotedIdentifier,
		PostgreSQLLexerUnicodeQuotedIdentifier,
		PostgreSQLLexerPLSQLVARIABLENAME,
		PostgreSQLLexerPLSQLIDENTIFIER,
		PostgreSQLLexerUnterminatedQuotedIdentifier:
		return true
	case PostgreSQLLexerABORT_P, // Unreserved Keywords
		PostgreSQLLexerABSOLUTE_P,
		PostgreSQLLexerACCESS,
		PostgreSQLLexerACTION,
		PostgreSQLLexerADD_P,
		PostgreSQLLexerADMIN,
		PostgreSQLLexerAFTER,
		PostgreSQLLexerAGGREGATE,
		PostgreSQLLexerALSO,
		PostgreSQLLexerALTER,
		PostgreSQLLexerALWAYS,
		PostgreSQLLexerASSERTION,
		PostgreSQLLexerASSIGNMENT,
		PostgreSQLLexerAT,
		PostgreSQLLexerATOMIC_P,
		PostgreSQLLexerATTACH,
		PostgreSQLLexerATTRIBUTE,
		PostgreSQLLexerBACKWARD,
		PostgreSQLLexerBEFORE,
		PostgreSQLLexerBEGIN_P,
		PostgreSQLLexerBY,
		PostgreSQLLexerCACHE,
		PostgreSQLLexerCALL,
		PostgreSQLLexerCALLED,
		PostgreSQLLexerCASCADE,
		PostgreSQLLexerCASCADED,
		PostgreSQLLexerCATALOG,
		PostgreSQLLexerCHAIN,
		PostgreSQLLexerCHARACTERISTICS,
		PostgreSQLLexerCHECKPOINT,
		PostgreSQLLexerCLASS,
		PostgreSQLLexerCLOSE,
		PostgreSQLLexerCLUSTER,
		PostgreSQLLexerCOLUMNS,
		PostgreSQLLexerCOMMENT,
		PostgreSQLLexerCOMMENTS,
		PostgreSQLLexerCOMMIT,
		PostgreSQLLexerCOMMITTED,
		PostgreSQLLexerCONFIGURATION,
		PostgreSQLLexerCONFLICT,
		PostgreSQLLexerCONNECTION,
		PostgreSQLLexerCONSTRAINTS,
		PostgreSQLLexerCONTENT_P,
		PostgreSQLLexerCONTINUE_P,
		PostgreSQLLexerCONVERSION_P,
		PostgreSQLLexerCOPY,
		PostgreSQLLexerCOST,
		PostgreSQLLexerCSV,
		PostgreSQLLexerCUBE,
		PostgreSQLLexerCURRENT_P,
		PostgreSQLLexerCURSOR,
		PostgreSQLLexerCYCLE,
		PostgreSQLLexerDATA_P,
		PostgreSQLLexerDATABASE,
		PostgreSQLLexerDAY_P,
		PostgreSQLLexerDEALLOCATE,
		PostgreSQLLexerDECLARE,
		PostgreSQLLexerDEFAULTS,
		PostgreSQLLexerDEFERRED,
		PostgreSQLLexerDEFINER,
		PostgreSQLLexerDELETE_P,
		PostgreSQLLexerDELIMITER,
		PostgreSQLLexerDELIMITERS,
		PostgreSQLLexerDEPENDS,
		PostgreSQLLexerDETACH,
		PostgreSQLLexerDICTIONARY,
		PostgreSQLLexerDISABLE_P,
		PostgreSQLLexerDISCARD,
		PostgreSQLLexerDOCUMENT_P,
		PostgreSQLLexerDOMAIN_P,
		PostgreSQLLexerDOUBLE_P,
		PostgreSQLLexerDROP,
		PostgreSQLLexerEACH,
		PostgreSQLLexerENABLE_P,
		PostgreSQLLexerENCODING,
		PostgreSQLLexerENCRYPTED,
		PostgreSQLLexerENUM_P,
		PostgreSQLLexerESCAPE,
		PostgreSQLLexerEVENT,
		PostgreSQLLexerEXCLUDE,
		PostgreSQLLexerEXCLUDING,
		PostgreSQLLexerEXCLUSIVE,
		PostgreSQLLexerEXECUTE,
		PostgreSQLLexerEXPLAIN,
		PostgreSQLLexerEXPRESSION,
		PostgreSQLLexerEXTENSION,
		PostgreSQLLexerEXTERNAL,
		PostgreSQLLexerFAMILY,
		PostgreSQLLexerFILTER,
		PostgreSQLLexerFIRST_P,
		PostgreSQLLexerFOLLOWING,
		PostgreSQLLexerFORCE,
		PostgreSQLLexerFORWARD,
		PostgreSQLLexerFUNCTION,
		PostgreSQLLexerFUNCTIONS,
		PostgreSQLLexerGENERATED,
		PostgreSQLLexerGLOBAL,
		PostgreSQLLexerGRANTED,
		PostgreSQLLexerGROUPS,
		PostgreSQLLexerHANDLER,
		PostgreSQLLexerHEADER_P,
		PostgreSQLLexerHOLD,
		PostgreSQLLexerHOUR_P,
		PostgreSQLLexerIDENTITY_P,
		PostgreSQLLexerIF_P,
		PostgreSQLLexerIMMEDIATE,
		PostgreSQLLexerIMMUTABLE,
		PostgreSQLLexerIMPLICIT_P,
		PostgreSQLLexerIMPORT_P,
		PostgreSQLLexerINCLUDE,
		PostgreSQLLexerINCLUDING,
		PostgreSQLLexerINCREMENT,
		PostgreSQLLexerINDEX,
		PostgreSQLLexerINDEXES,
		PostgreSQLLexerINHERIT,
		PostgreSQLLexerINHERITS,
		PostgreSQLLexerINLINE_P,
		PostgreSQLLexerINPUT_P,
		PostgreSQLLexerINSENSITIVE,
		PostgreSQLLexerINSERT,
		PostgreSQLLexerINSTEAD,
		PostgreSQLLexerINVOKER,
		PostgreSQLLexerISOLATION,
		PostgreSQLLexerKEY,
		PostgreSQLLexerLABEL,
		PostgreSQLLexerLANGUAGE,
		PostgreSQLLexerLARGE_P,
		PostgreSQLLexerLAST_P,
		PostgreSQLLexerLEAKPROOF,
		PostgreSQLLexerLEVEL,
		PostgreSQLLexerLISTEN,
		PostgreSQLLexerLOAD,
		PostgreSQLLexerLOCAL,
		PostgreSQLLexerLOCATION,
		PostgreSQLLexerLOCK_P,
		PostgreSQLLexerLOCKED,
		PostgreSQLLexerLOGGED,
		PostgreSQLLexerMAPPING,
		PostgreSQLLexerMATCH,
		PostgreSQLLexerMATERIALIZED,
		PostgreSQLLexerMAXVALUE,
		PostgreSQLLexerMETHOD,
		PostgreSQLLexerMINUTE_P,
		PostgreSQLLexerMINVALUE,
		PostgreSQLLexerMODE,
		PostgreSQLLexerMONTH_P,
		PostgreSQLLexerMOVE,
		PostgreSQLLexerNAME_P,
		PostgreSQLLexerNAMES,
		PostgreSQLLexerNEW,
		PostgreSQLLexerNEXT,
		PostgreSQLLexerNFC,
		PostgreSQLLexerNFD,
		PostgreSQLLexerNFKC,
		PostgreSQLLexerNFKD,
		PostgreSQLLexerNO,
		PostgreSQLLexerNORMALIZED,
		PostgreSQLLexerNOTHING,
		PostgreSQLLexerNOTIFY,
		PostgreSQLLexerNOWAIT,
		PostgreSQLLexerNULLS_P,
		PostgreSQLLexerOBJECT_P,
		PostgreSQLLexerOF,
		PostgreSQLLexerOFF,
		PostgreSQLLexerOIDS,
		PostgreSQLLexerOLD,
		PostgreSQLLexerOPERATOR,
		PostgreSQLLexerOPTION,
		PostgreSQLLexerOPTIONS,
		PostgreSQLLexerORDINALITY,
		PostgreSQLLexerOTHERS,
		PostgreSQLLexerOVER,
		PostgreSQLLexerOVERRIDING,
		PostgreSQLLexerOWNED,
		PostgreSQLLexerOWNER,
		PostgreSQLLexerPARALLEL,
		PostgreSQLLexerPARSER,
		PostgreSQLLexerPARTIAL,
		PostgreSQLLexerPARTITION,
		PostgreSQLLexerPASSING,
		PostgreSQLLexerPASSWORD,
		PostgreSQLLexerPLANS,
		PostgreSQLLexerPOLICY,
		PostgreSQLLexerPRECEDING,
		PostgreSQLLexerPREPARE,
		PostgreSQLLexerPREPARED,
		PostgreSQLLexerPRESERVE,
		PostgreSQLLexerPRIOR,
		PostgreSQLLexerPRIVILEGES,
		PostgreSQLLexerPROCEDURAL,
		PostgreSQLLexerPROCEDURE,
		PostgreSQLLexerPROCEDURES,
		PostgreSQLLexerPROGRAM,
		PostgreSQLLexerPUBLICATION,
		PostgreSQLLexerQUOTE,
		PostgreSQLLexerRANGE,
		PostgreSQLLexerREAD,
		PostgreSQLLexerREASSIGN,
		PostgreSQLLexerRECHECK,
		PostgreSQLLexerRECURSIVE,
		PostgreSQLLexerREF,
		PostgreSQLLexerREFERENCING,
		PostgreSQLLexerREFRESH,
		PostgreSQLLexerREINDEX,
		PostgreSQLLexerRELATIVE_P,
		PostgreSQLLexerRELEASE,
		PostgreSQLLexerRENAME,
		PostgreSQLLexerREPEATABLE,
		PostgreSQLLexerREPLICA,
		PostgreSQLLexerRESET,
		PostgreSQLLexerRESTART,
		PostgreSQLLexerRESTRICT,
		PostgreSQLLexerRETURNS,
		PostgreSQLLexerREVOKE,
		PostgreSQLLexerROLE,
		PostgreSQLLexerROLLBACK,
		PostgreSQLLexerROLLUP,
		PostgreSQLLexerROUTINE,
		PostgreSQLLexerROUTINES,
		PostgreSQLLexerROWS,
		PostgreSQLLexerRULE,
		PostgreSQLLexerSAVEPOINT,
		PostgreSQLLexerSCHEMA,
		PostgreSQLLexerSCHEMAS,
		PostgreSQLLexerSCROLL,
		PostgreSQLLexerSEARCH,
		PostgreSQLLexerSECOND_P,
		PostgreSQLLexerSECURITY,
		PostgreSQLLexerSEQUENCE,
		PostgreSQLLexerSEQUENCES,
		PostgreSQLLexerSERIALIZABLE,
		PostgreSQLLexerSERVER,
		PostgreSQLLexerSESSION,
		PostgreSQLLexerSET,
		PostgreSQLLexerSETS,
		PostgreSQLLexerSHARE,
		PostgreSQLLexerSHOW,
		PostgreSQLLexerSIMPLE,
		PostgreSQLLexerSKIP_P,
		PostgreSQLLexerSNAPSHOT,
		PostgreSQLLexerSQL_P,
		PostgreSQLLexerSTABLE,
		PostgreSQLLexerSTANDALONE_P,
		PostgreSQLLexerSTART,
		PostgreSQLLexerSTATEMENT,
		PostgreSQLLexerSTATISTICS,
		PostgreSQLLexerSTDIN,
		PostgreSQLLexerSTDOUT,
		PostgreSQLLexerSTORAGE,
		PostgreSQLLexerSTORED,
		PostgreSQLLexerSTRICT_P,
		PostgreSQLLexerSTRIP_P,
		PostgreSQLLexerSUBSCRIPTION,
		PostgreSQLLexerSUPPORT,
		PostgreSQLLexerSYSID,
		PostgreSQLLexerSYSTEM_P,
		PostgreSQLLexerTABLES,
		PostgreSQLLexerTABLESPACE,
		PostgreSQLLexerTEMP,
		PostgreSQLLexerTEMPLATE,
		PostgreSQLLexerTEMPORARY,
		PostgreSQLLexerTEXT_P,
		PostgreSQLLexerTIES,
		PostgreSQLLexerTRANSACTION,
		PostgreSQLLexerTRANSFORM,
		PostgreSQLLexerTRIGGER,
		PostgreSQLLexerTRUNCATE,
		PostgreSQLLexerTRUSTED,
		PostgreSQLLexerTYPE_P,
		PostgreSQLLexerTYPES_P,
		PostgreSQLLexerUESCAPE,
		PostgreSQLLexerUNBOUNDED,
		PostgreSQLLexerUNCOMMITTED,
		PostgreSQLLexerUNENCRYPTED,
		PostgreSQLLexerUNKNOWN,
		PostgreSQLLexerUNLISTEN,
		PostgreSQLLexerUNLOGGED,
		PostgreSQLLexerUNTIL,
		PostgreSQLLexerUPDATE,
		PostgreSQLLexerVACUUM,
		PostgreSQLLexerVALID,
		PostgreSQLLexerVALIDATE,
		PostgreSQLLexerVALIDATOR,
		PostgreSQLLexerVALUE_P,
		PostgreSQLLexerVARYING,
		PostgreSQLLexerVERSION_P,
		PostgreSQLLexerVIEW,
		PostgreSQLLexerVIEWS,
		PostgreSQLLexerVOLATILE,
		PostgreSQLLexerWHITESPACE_P,
		PostgreSQLLexerWITHIN,
		PostgreSQLLexerWITHOUT,
		PostgreSQLLexerWORK,
		PostgreSQLLexerWRAPPER,
		PostgreSQLLexerWRITE,
		PostgreSQLLexerXML_P,
		PostgreSQLLexerYEAR_P,
		PostgreSQLLexerYES_P,
		PostgreSQLLexerZONE:
		return true
	case PostgreSQLLexerALIAS, // plsql unreserved keywords
		PostgreSQLLexerAND,
		PostgreSQLLexerARRAY,
		PostgreSQLLexerASSERT,
		PostgreSQLLexerCOLLATE,
		PostgreSQLLexerCOLUMN,
		PostgreSQLLexerCONSTANT,
		PostgreSQLLexerCONSTRAINT,
		PostgreSQLLexerDEBUG,
		PostgreSQLLexerDEFAULT,
		PostgreSQLLexerDIAGNOSTICS,
		PostgreSQLLexerDO,
		PostgreSQLLexerDUMP,
		PostgreSQLLexerELSIF,
		PostgreSQLLexerERROR,
		PostgreSQLLexerEXCEPTION,
		PostgreSQLLexerEXIT,
		PostgreSQLLexerFETCH,
		PostgreSQLLexerGET,
		PostgreSQLLexerINFO,
		PostgreSQLLexerIS,
		PostgreSQLLexerNOTICE,
		PostgreSQLLexerOPEN,
		PostgreSQLLexerPERFORM,
		PostgreSQLLexerPRINT_STRICT_PARAMS,
		PostgreSQLLexerQUERY,
		PostgreSQLLexerRAISE,
		PostgreSQLLexerRETURN,
		PostgreSQLLexerSLICE,
		PostgreSQLLexerSQLSTATE,
		PostgreSQLLexerSTACKED,
		PostgreSQLLexerTABLE,
		PostgreSQLLexerUSE_COLUMN,
		PostgreSQLLexerUSE_VARIABLE,
		PostgreSQLLexerVARIABLE_CONFLICT,
		PostgreSQLLexerWARNING,
		PostgreSQLLexerOUTER_P:
		return true
	}

	symbol := receiver.GetSymbolicNames()[tokenType]
	symbol = strings.TrimSuffix(symbol, "_SYMBOL")
	if len(symbol) > 0 && !receiver.IsReservedKeyword(symbol) {
		return true
	}
	return false
}

func (receiver *PostgreSQLLexerBase) IsReservedKeyword(s string) bool {
	receiver.initReservedKeywordMap()

	_, exists := receiver.reservedKeywordMap[s]
	return exists
}

func (receiver *PostgreSQLLexerBase) initReservedKeywordMap() {
	if receiver.reservedKeywordMap != nil {
		return
	}

	receiver.reservedKeywordMap = make(map[string]bool)

	for _, keyword := range Keywords {
		if keyword.Reserved {
			receiver.reservedKeywordMap[keyword.Keyword] = true
		}
	}
}
