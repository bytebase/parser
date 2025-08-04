package redshift

import (
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

type RedshiftLexerBase struct {
	*antlr.BaseLexer

	stack              StringStack
	reservedKeywordMap map[string]bool
}

func (receiver *RedshiftLexerBase) pushTag() {
	receiver.stack.Push(receiver.GetText())
}

func (receiver *RedshiftLexerBase) isTag() bool {
	if receiver.stack.IsEmpty() {
		return false
	}
	return receiver.GetText() == receiver.stack.PeekOrEmpty()
}

func (receiver *RedshiftLexerBase) popTag() {
	_, _ = receiver.stack.Pop()
}

func (receiver *RedshiftLexerBase) checkLA(c int) bool {
	return receiver.GetInputStream().LA(1) != c
}

func (receiver *RedshiftLexerBase) charIsLetter() bool {
	c := receiver.GetInputStream().LA(-1)
	return unicode.IsLetter(rune(c))
}

func (receiver *RedshiftLexerBase) HandleNumericFail() {
	index := receiver.GetInputStream().Index() - 2
	receiver.GetInputStream().Seek(index)
	receiver.SetType(RedshiftLexerIntegral)
}

func (receiver *RedshiftLexerBase) HandleLessLessGreaterGreater() {
	if receiver.GetText() == "<<" {
		receiver.SetType(RedshiftLexerLESS_LESS)
	}
	if receiver.GetText() == ">>" {
		receiver.SetType(RedshiftLexerGREATER_GREATER)
	}
}

func (receiver *RedshiftLexerBase) UnterminatedBlockCommentDebugAssert() {
	//Debug.Assert(InputStream.LA(1) == -1 /*EOF*/);
}

func (receiver *RedshiftLexerBase) CheckIfUtf32Letter() bool {
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

func (receiver *RedshiftLexerBase) IsIdentifier(tokenType int) bool {
	if tokenType == antlr.TokenEOF {
		return false
	}

	switch tokenType {
	case RedshiftLexerIdentifier,
		RedshiftLexerQuotedIdentifier,
		RedshiftLexerUnicodeQuotedIdentifier,
		RedshiftLexerPLSQLVARIABLENAME,
		RedshiftLexerPLSQLIDENTIFIER,
		RedshiftLexerUnterminatedQuotedIdentifier:
		return true
	case RedshiftLexerABORT_P, // Unreserved Keywords
		RedshiftLexerABSOLUTE_P,
		RedshiftLexerACCESS,
		RedshiftLexerACTION,
		RedshiftLexerADD_P,
		RedshiftLexerADMIN,
		RedshiftLexerAFTER,
		RedshiftLexerAGGREGATE,
		RedshiftLexerALSO,
		RedshiftLexerALTER,
		RedshiftLexerALWAYS,
		RedshiftLexerASSERTION,
		RedshiftLexerASSIGNMENT,
		RedshiftLexerAT,
		RedshiftLexerATOMIC_P,
		RedshiftLexerATTACH,
		RedshiftLexerATTRIBUTE,
		RedshiftLexerBACKWARD,
		RedshiftLexerBEFORE,
		RedshiftLexerBEGIN_P,
		RedshiftLexerBY,
		RedshiftLexerCACHE,
		RedshiftLexerCALL,
		RedshiftLexerCALLED,
		RedshiftLexerCASCADE,
		RedshiftLexerCASCADED,
		RedshiftLexerCATALOG,
		RedshiftLexerCHAIN,
		RedshiftLexerCHARACTERISTICS,
		RedshiftLexerCHECKPOINT,
		RedshiftLexerCLASS,
		RedshiftLexerCLOSE,
		RedshiftLexerCLUSTER,
		RedshiftLexerCOLUMNS,
		RedshiftLexerCOMMENT,
		RedshiftLexerCOMMENTS,
		RedshiftLexerCOMMIT,
		RedshiftLexerCOMMITTED,
		RedshiftLexerCONFIGURATION,
		RedshiftLexerCONFLICT,
		RedshiftLexerCONNECTION,
		RedshiftLexerCONSTRAINTS,
		RedshiftLexerCONTENT_P,
		RedshiftLexerCONTINUE_P,
		RedshiftLexerCONVERSION_P,
		RedshiftLexerCOPY,
		RedshiftLexerCOST,
		RedshiftLexerCSV,
		RedshiftLexerCUBE,
		RedshiftLexerCURRENT_P,
		RedshiftLexerCURSOR,
		RedshiftLexerCYCLE,
		RedshiftLexerDATA_P,
		RedshiftLexerDATABASE,
		RedshiftLexerDAY_P,
		RedshiftLexerDEALLOCATE,
		RedshiftLexerDECLARE,
		RedshiftLexerDEFAULTS,
		RedshiftLexerDEFERRED,
		RedshiftLexerDEFINER,
		RedshiftLexerDELETE_P,
		RedshiftLexerDELIMITER,
		RedshiftLexerDELIMITERS,
		RedshiftLexerDEPENDS,
		RedshiftLexerDETACH,
		RedshiftLexerDICTIONARY,
		RedshiftLexerDISABLE_P,
		RedshiftLexerDISCARD,
		RedshiftLexerDOCUMENT_P,
		RedshiftLexerDOMAIN_P,
		RedshiftLexerDOUBLE_P,
		RedshiftLexerDROP,
		RedshiftLexerEACH,
		RedshiftLexerENABLE_P,
		RedshiftLexerENCODING,
		RedshiftLexerENCRYPTED,
		RedshiftLexerENUM_P,
		RedshiftLexerESCAPE,
		RedshiftLexerEVENT,
		RedshiftLexerEXCLUDE,
		RedshiftLexerEXCLUDING,
		RedshiftLexerEXCLUSIVE,
		RedshiftLexerEXECUTE,
		RedshiftLexerEXPLAIN,
		RedshiftLexerEXPRESSION,
		RedshiftLexerEXTENSION,
		RedshiftLexerEXTERNAL,
		RedshiftLexerFAMILY,
		RedshiftLexerFILTER,
		RedshiftLexerFIRST_P,
		RedshiftLexerFOLLOWING,
		RedshiftLexerFORCE,
		RedshiftLexerFORWARD,
		RedshiftLexerFUNCTION,
		RedshiftLexerFUNCTIONS,
		RedshiftLexerGENERATED,
		RedshiftLexerGLOBAL,
		RedshiftLexerGRANTED,
		RedshiftLexerGROUPS,
		RedshiftLexerHANDLER,
		RedshiftLexerHEADER_P,
		RedshiftLexerHOLD,
		RedshiftLexerHOUR_P,
		RedshiftLexerIDENTITY_P,
		RedshiftLexerIF_P,
		RedshiftLexerIMMEDIATE,
		RedshiftLexerIMMUTABLE,
		RedshiftLexerIMPLICIT_P,
		RedshiftLexerIMPORT_P,
		RedshiftLexerINCLUDE,
		RedshiftLexerINCLUDING,
		RedshiftLexerINCREMENT,
		RedshiftLexerINDEX,
		RedshiftLexerINDEXES,
		RedshiftLexerINHERIT,
		RedshiftLexerINHERITS,
		RedshiftLexerINLINE_P,
		RedshiftLexerINPUT_P,
		RedshiftLexerINSENSITIVE,
		RedshiftLexerINSERT,
		RedshiftLexerINSTEAD,
		RedshiftLexerINVOKER,
		RedshiftLexerISOLATION,
		RedshiftLexerKEY,
		RedshiftLexerLABEL,
		RedshiftLexerLANGUAGE,
		RedshiftLexerLARGE_P,
		RedshiftLexerLAST_P,
		RedshiftLexerLEAKPROOF,
		RedshiftLexerLEVEL,
		RedshiftLexerLISTEN,
		RedshiftLexerLOAD,
		RedshiftLexerLOCAL,
		RedshiftLexerLOCATION,
		RedshiftLexerLOCK_P,
		RedshiftLexerLOCKED,
		RedshiftLexerLOGGED,
		RedshiftLexerMAPPING,
		RedshiftLexerMATCH,
		RedshiftLexerMATERIALIZED,
		RedshiftLexerMAXVALUE,
		RedshiftLexerMETHOD,
		RedshiftLexerMINUTE_P,
		RedshiftLexerMINVALUE,
		RedshiftLexerMODE,
		RedshiftLexerMONTH_P,
		RedshiftLexerMOVE,
		RedshiftLexerNAME_P,
		RedshiftLexerNAMES,
		RedshiftLexerNEW,
		RedshiftLexerNEXT,
		RedshiftLexerNFC,
		RedshiftLexerNFD,
		RedshiftLexerNFKC,
		RedshiftLexerNFKD,
		RedshiftLexerNO,
		RedshiftLexerNORMALIZED,
		RedshiftLexerNOTHING,
		RedshiftLexerNOTIFY,
		RedshiftLexerNOWAIT,
		RedshiftLexerNULLS_P,
		RedshiftLexerOBJECT_P,
		RedshiftLexerOF,
		RedshiftLexerOFF,
		RedshiftLexerOIDS,
		RedshiftLexerOLD,
		RedshiftLexerOPERATOR,
		RedshiftLexerOPTION,
		RedshiftLexerOPTIONS,
		RedshiftLexerORDINALITY,
		RedshiftLexerOTHERS,
		RedshiftLexerOVER,
		RedshiftLexerOVERRIDING,
		RedshiftLexerOWNED,
		RedshiftLexerOWNER,
		RedshiftLexerPARALLEL,
		RedshiftLexerPARSER,
		RedshiftLexerPARTIAL,
		RedshiftLexerPARTITION,
		RedshiftLexerPASSING,
		RedshiftLexerPASSWORD,
		RedshiftLexerPLANS,
		RedshiftLexerPOLICY,
		RedshiftLexerPRECEDING,
		RedshiftLexerPREPARE,
		RedshiftLexerPREPARED,
		RedshiftLexerPRESERVE,
		RedshiftLexerPRIOR,
		RedshiftLexerPRIVILEGES,
		RedshiftLexerPROCEDURAL,
		RedshiftLexerPROCEDURE,
		RedshiftLexerPROCEDURES,
		RedshiftLexerPROGRAM,
		RedshiftLexerPUBLICATION,
		RedshiftLexerQUOTE,
		RedshiftLexerRANGE,
		RedshiftLexerREAD,
		RedshiftLexerREASSIGN,
		RedshiftLexerRECHECK,
		RedshiftLexerRECURSIVE,
		RedshiftLexerREF,
		RedshiftLexerREFERENCING,
		RedshiftLexerREFRESH,
		RedshiftLexerREINDEX,
		RedshiftLexerRELATIVE_P,
		RedshiftLexerRELEASE,
		RedshiftLexerRENAME,
		RedshiftLexerREPEATABLE,
		RedshiftLexerREPLICA,
		RedshiftLexerRESET,
		RedshiftLexerRESTART,
		RedshiftLexerRESTRICT,
		RedshiftLexerRETURNS,
		RedshiftLexerREVOKE,
		RedshiftLexerROLE,
		RedshiftLexerROLLBACK,
		RedshiftLexerROLLUP,
		RedshiftLexerROUTINE,
		RedshiftLexerROUTINES,
		RedshiftLexerROWS,
		RedshiftLexerRULE,
		RedshiftLexerSAVEPOINT,
		RedshiftLexerSCHEMA,
		RedshiftLexerSCHEMAS,
		RedshiftLexerSCROLL,
		RedshiftLexerSEARCH,
		RedshiftLexerSECOND_P,
		RedshiftLexerSECURITY,
		RedshiftLexerSEQUENCE,
		RedshiftLexerSEQUENCES,
		RedshiftLexerSERIALIZABLE,
		RedshiftLexerSERVER,
		RedshiftLexerSESSION,
		RedshiftLexerSET,
		RedshiftLexerSETS,
		RedshiftLexerSHARE,
		RedshiftLexerSHOW,
		RedshiftLexerSIMPLE,
		RedshiftLexerSKIP_P,
		RedshiftLexerSNAPSHOT,
		RedshiftLexerSQL_P,
		RedshiftLexerSTABLE,
		RedshiftLexerSTANDALONE_P,
		RedshiftLexerSTART,
		RedshiftLexerSTATEMENT,
		RedshiftLexerSTATISTICS,
		RedshiftLexerSTDIN,
		RedshiftLexerSTDOUT,
		RedshiftLexerSTORAGE,
		RedshiftLexerSTORED,
		RedshiftLexerSTRICT_P,
		RedshiftLexerSTRIP_P,
		RedshiftLexerSUBSCRIPTION,
		RedshiftLexerSUPPORT,
		RedshiftLexerSYSID,
		RedshiftLexerSYSTEM_P,
		RedshiftLexerTABLES,
		RedshiftLexerTABLESPACE,
		RedshiftLexerTEMP,
		RedshiftLexerTEMPLATE,
		RedshiftLexerTEMPORARY,
		RedshiftLexerTEXT_P,
		RedshiftLexerTIES,
		RedshiftLexerTRANSACTION,
		RedshiftLexerTRANSFORM,
		RedshiftLexerTRIGGER,
		RedshiftLexerTRUNCATE,
		RedshiftLexerTRUSTED,
		RedshiftLexerTYPE_P,
		RedshiftLexerTYPES_P,
		RedshiftLexerUESCAPE,
		RedshiftLexerUNBOUNDED,
		RedshiftLexerUNCOMMITTED,
		RedshiftLexerUNENCRYPTED,
		RedshiftLexerUNKNOWN,
		RedshiftLexerUNLISTEN,
		RedshiftLexerUNLOGGED,
		RedshiftLexerUNTIL,
		RedshiftLexerUPDATE,
		RedshiftLexerVACUUM,
		RedshiftLexerVALID,
		RedshiftLexerVALIDATE,
		RedshiftLexerVALIDATOR,
		RedshiftLexerVALUE_P,
		RedshiftLexerVARYING,
		RedshiftLexerVERSION_P,
		RedshiftLexerVIEW,
		RedshiftLexerVIEWS,
		RedshiftLexerVOLATILE,
		RedshiftLexerWHITESPACE_P,
		RedshiftLexerWITHIN,
		RedshiftLexerWITHOUT,
		RedshiftLexerWORK,
		RedshiftLexerWRAPPER,
		RedshiftLexerWRITE,
		RedshiftLexerXML_P,
		RedshiftLexerYEAR_P,
		RedshiftLexerYES_P,
		RedshiftLexerZONE:
		return true
	case RedshiftLexerALIAS, // plsql unreserved keywords
		RedshiftLexerAND,
		RedshiftLexerARRAY,
		RedshiftLexerASSERT,
		RedshiftLexerCOLLATE,
		RedshiftLexerCOLUMN,
		RedshiftLexerCONSTANT,
		RedshiftLexerCONSTRAINT,
		RedshiftLexerDEBUG,
		RedshiftLexerDEFAULT,
		RedshiftLexerDIAGNOSTICS,
		RedshiftLexerDO,
		RedshiftLexerDUMP,
		RedshiftLexerELSIF,
		RedshiftLexerERROR,
		RedshiftLexerEXCEPTION,
		RedshiftLexerEXIT,
		RedshiftLexerFETCH,
		RedshiftLexerGET,
		RedshiftLexerINFO,
		RedshiftLexerIS,
		RedshiftLexerNOTICE,
		RedshiftLexerOPEN,
		RedshiftLexerPERFORM,
		RedshiftLexerPRINT_STRICT_PARAMS,
		RedshiftLexerQUERY,
		RedshiftLexerRAISE,
		RedshiftLexerRETURN,
		RedshiftLexerSLICE,
		RedshiftLexerSQLSTATE,
		RedshiftLexerSTACKED,
		RedshiftLexerTABLE,
		RedshiftLexerUSE_COLUMN,
		RedshiftLexerUSE_VARIABLE,
		RedshiftLexerVARIABLE_CONFLICT,
		RedshiftLexerWARNING,
		RedshiftLexerOUTER_P:
		return true
	}

	symbol := receiver.GetSymbolicNames()[tokenType]
	symbol = strings.TrimSuffix(symbol, "_SYMBOL")
	if len(symbol) > 0 && !receiver.IsReservedKeyword(symbol) {
		return true
	}
	return false
}

func (receiver *RedshiftLexerBase) IsReservedKeyword(s string) bool {
	receiver.initReservedKeywordMap()

	_, exists := receiver.reservedKeywordMap[s]
	return exists
}

func (receiver *RedshiftLexerBase) initReservedKeywordMap() {
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
