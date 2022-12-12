package logic

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func Test_Split(t *testing.T) {

	f, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	tests := []struct {
		name    string
		in      []byte
		want    [][]string
		wantErr bool
	}{
		{
			name: "Small Test",
			in: []byte(
				`vJrwpWtwJgWrhcsFMMfFFhFp
			jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
			PmmdzqPrVvPwwTWBwg`),
			want:    [][]string{{"vJrwpWtwJgWr", "jqHRNqRjqzjGDLGL", "PmmdzqPrV"}, {"hcsFMMfFFhFp", "rsFMfFZSrLrFZsSL", "vPwwTWBwg"}},
			wantErr: false,
		},
		{
			name:    "Real Test",
			in:      f,
			want:    [][]string{{"jNNBMTNzvT", "VCwnVRCGHHJT", "wFJZTbRcnJCbpwpFccZC", "wDWgDfWNTvwvgFfWf", "nLnmLSnmM", "MsmsbLvtzMj", "WTqSCqWSWqSgVZqJ", "PtLGLGdd", "jFvLPPlLjvfjjffsclvP", "lLSSrfmddlNpnmLdfSPddDd", "MJMnhzBJVJwvGsMRhRh", "vzvBWGJMVGwWGGhGqVBnGzVG", "qGWLgfbWhqpLgZb", "FQwPwNPCVFQQdNcFHNcww", "wQHnwnQFPNPdjnnP", "mpVpmtPhVtP", "vzFMlMzvvR", "srqzvRRqdvzFFrDnvqldF", "CQJJLZCJLq", "RfRFFszTtzbWTFnRtFrRrW", "znsjWzstnWjbDnDbT", "jRrzzrPjL", "CTVnTlZWHTcqZBcTTTHqW", "qlWFWWHTZTlqTFWZHtBVFdlT", "cWjbcjnssMmRPWbGsWcMbf", "vqpvFHNvhvJzd", "JzqVVVVpzhVJqzhSzJhhDV", "NznNGQgS", "jjjCLwdvLcrhq", "qrrvwHqjtL", "LGPtFtTfTfTfLrrWTSWGPrr", "NhHpVpqNN", "ccBMJVDNCvvhDNpvNDpnNCVC", "dRgRhVLRl", "WJmwvvvjQvWwFBBvFJPHpwv", "FFHJjmjjjPMvfHfw", "NDTRFSnN", "JjLZWqwjHvfrvmHHvjLWWL", "mwHvHmpvwpLmjnvpvLvv", "GGhZQMsmGRfMwfp", "tFvFSTSDnDTTLt", "vzDtzMVtnLTTMTSjSDHt", "MfJtWTTMJfzB", "FvlgSHSLvmCvZjgcC", "HLhclRlHRmHvmLmclLZFSF", "scsswLQcGmQm", "bdnPSMMMClbBBPCdDDnDVMD", "bDPTlSTPPbSPjnPSb", "tFFJjMnFhdcMMJlWtdnlFc", "QRgqwwvs", "gqfvvnpfHwRmH", "fccTzVVcfSm", "wLsWfppsjBrnnwjCBZnZ", "pZrNCWLwLZLjwwW", "rZnVVjVS", "dLvQLBLFddvBLzpGmddQC", "LLCCWjGBFPmCWdmmPWLdWBpL", "ddfnQRbpldRlR", "BWBVDZWTBTDPvVTZVDBDNMbh", "NPZPDMWZMbZrBVMDDWNmBhMV", "bvPSBttgG", "NJspjJLdprzHgjrLzLNrnH", "HdRJsJrdHT", "sNVvSdTstDCtdzdzSCwTzC", "rgWlrnLLbqb", "BmrnrFPnnBLlq", "mJMqlVlttQlFVmzFQMQMbQM", "GGZgcfjrvCCZvgCZqvPH", "GZppfGHGrgGprZh", "WDHHLtRBHgDnVrWFVFB", "JlqmMpCMCGM", "cJJqSplpfSGccJGt", "tlFwbWtQFLJhlBFlWPbwm", "GCMBDDCvRpDVpRpH", "MDDvCVGzBGfpzBSNMSRR", "fzFzFHBfnvpHFbnzbHfBHZ", "sLMNSccQLMSrmLcshL", "cwLMwwNhQhsWS", "NWlqqhNNnGtNvvW", "mppCjppMrDTSgDppCDT", "pgzzTFFFrj", "ncQrhQjqjVQh", "RllWFLRf", "vFLLmfNRFSNDfS", "mmGrgwmGD", "nDsqfPCHnpntDssfJPDP", "DfPCJLpbl", "dJDNbRhNbJdhq", "LMHrttLFcMmcMFLmtG", "czZMzrzVZzHZ", "sbsJgbsm", "LLLFBzGMLjzzFtVFwww", "GllfptwjLjlGFVFjGzpFF", "SccPbfbncpcfsjbRjM", "vvvdgVhlmlwlgZLgm", "NbLNDdbLVmqqm", "HcgDBJSHTC", "VppGMwFwllwwbZZ", "wZlhtPGPLFVFlGZbVtV", "HPMMnhBHlMnMBPBHJHPWfdnB", "SrzrSScrctrwVzCSC", "CswstFRNpcwVNRrVVSVwpw", "mlFMtqjvMdqjmM", "HcNgcPLcHLwtcf", "LGGNhbNtNLDfcgcwfbgthfw", "NbvbBGNvhNhnhpbgpGfBvNg", "qqFSSqWrtSScc", "SBDSBZZZ", "qTdHSLSFRZRHHZV", "tDvflvrGttDCjlQfmCGvtCft", "tQlDCDtrfrtlGrjlQB", "MTwvsQJMvvHwVMMJMQNNJRP", "zhcmnRZmqm", "ZbmhZcdZznZfndt", "ZhhVqQTvZvVhSmQZ", "dJWHwndMHswsw", "dnwwNJnjBjwwdddnjndfW", "trlrlrZzsjRj", "QqLvvLQmLfQTLbLTfHmqHHLq", "DPDvqLqffqf", "gsnVdcBc", "RQZLRZlZDRJDCzCjblJSDjQZ", "pDlQbSlzpQ", "zgqtCtJltdGttJgVGPPJCtJv", "rRcwBnnpcmBq", "mwRqrqrmWrZnfWmBmnR", "jlzHllmPnpHlH", "bQrhDbcLgsQrtd", "QhGdGcLGwwwtcD", "qbpqvWFHbFHHsWwPqPpsVW", "MCgSSJMCzt", "QCRMRSnNtRS", "ZnZrTfsWWv", "pLGBLBgLCpgGpbdLbgMClJ", "dpJPldgMMbglFd", "CCZCQzwwdmMGDWMmhCM", "jfStHtcjqDPbPtvqv", "fcqqPPPtPSfRVBBRPRPbv", "VVlDNDgppgtNltlrJP", "LmGRmfmwCq", "cHRcCFfmfmGLJLGFcJLFwf", "vSSdnBVp", "mLHzqtGtNfGHHFNH", "GHqHHGzNDNFCfqllCFq", "CMWcwGTrvzDWzrDccDCGzTTT", "nssRttdsnhsdbFhtVFhN", "PnVPVPLFt", "tlQlqlJCCJWgF", "SDbwShHbBZZbhcBdDBDST", "BTbbScwRhwZL", "VDfPVHfDVMMfHSPSMTVfqg", "bWGGHCplWnJNpJCCnln", "HpGCnZGllzlzJWjnCbbnCnR", "fQqcfqfS", "NcCLCtPCPMtNBwdthpLhPw", "twMPBtCpLBCtwMPpJcMB", "vQwVQQVqcJvVJvCpQBCLp", "hszPRDTVD", "TtDDZsszlPllhPzmPVGssTP", "sRmGqqzzz", "cffSwNDddTdf", "GGcTwHwHv", "JSdjLJMBdMSrfwwLp", "VsRtVllstHHNllsQsHQ", "tGVGCTTTtTFVbsTTNTHsTT", "MhjTJjlS", "RHttqbNGGmbbLmLmdqbgn", "gCNmDHHRCLHn", "BvzpbBwBms", "TnNRdtlRlCd", "jltNtRCPSlPNtc", "hflJphNDmmbpfnfpl", "FgTFWqTBfFPfFqQqT", "BSSRTGWBgFggBWgBGStTG", "mDDFjjFmV", "hhRvNNCvdNMRLzhSzpp", "WCdMzdMzhMdvRvLhCRLPv", "cScMdhsDhDDdv", "NPWjJqrjJWwrSFWRJlrlNWJ", "NWNQwNPj", "fVVCVccppZMZMMC", "QrmmPhDqPhsPRhrlbgRDbbPD", "DgbPLbPsQsrbsqLbgllsQQ", "ZlsmlrZZJcQmhBhlNrsrJ", "fMFvdfVTTdjWTTTfvSjVjpj", "cdvfSWjfjpdVqSwSvMd", "sWVVmDJsNWNjcSNJZcNcZWW", "nqPRwQRgpQRPQQgMQgQLQqB", "qgrPMLLPp", "VcGjcCHcVHPrGnjQ", "bWWhfzJhvZWJzNpm", "MhJMhJbMmSZbJZ", "FCvpgDsZNsCbvvvp", "fhfWdPhhhhRQQqzdLSRHRH", "qndhdQWndftfhWStfLQQfW", "mqTCNhTNmGTLwLN", "cjsMQppsFnFslnRQQ", "stWFjcFFsJlQcQt", "fmsffcqhmqPsnTCnCc", "vHbgLRvvv", "WltLWvWl", "CtwjffWrdznRtzC", "cZFFbgPgJZDc", "SpNFJBJccNgDcTJJT", "qWzgNFqzqHNTBzFN", "rfSJRJnhhnJDjrfvRSt", "cffRvfnnVRj", "VcHhVrVCQQWh", "dDSCGSsSblwDdmLqv", "tCCwGmSqbqtws", "fvvTcWzGcCJrJGJvv", "SgFhVgllLgjLgwlwljFqVFSF", "LNqFnlDgVqllwLFLnVSgLF", "FczpzmSjVVpSQrzzcRpRcr", "TvbGGbNfGCBB", "shgdNTGgvNsfls", "JHLPLTsSllgSSPPSPLTRT", "mhFpnpFw", "prpWcvnmhmpccBB", "GcMcjDbDMM", "rLCrwNJCnwrZNLWQQwz", "zhnrnzdrCwLJCwzwC", "cnwlFrdMsggblgsrMbncwrs", "vvVVHSpQvvRQDJ", "hCmJqVvMSpqqChCJHD", "sfDNqLNpqpzCzLsDqzbCVW", "wZrlwJwFJGlFM", "wrQrgnPwM", "hDcwwGWh", "NQmLbNSZHQSHSNpb", "LNVtLSSHLSgLNr", "CRDfCbfjcnRCBVfjVMfMj", "HdTFLGsdLrzN", "JLGmmHrwJGw", "VShGpPbWj", "HqfqDfDFFJDvZRJvqZRRqHZd", "FddsFmJjj", "NRFFLtFtqF", "ljMfMBDlJHgBVHgVflf", "PlzDzjlj", "dtzZZbctPzwdlzRwlcdfRgt", "nDjHnNjjjMnfjNTM", "CDpVCfrjLnnf", "sssZSZtDfHbbdt", "jwWgmWlNQNLlcjWhgQlrQQWr", "ggLgrqljLl", "cSGBFsFcSRZSQGsgBNg", "vnCTLlrpPlH", "rJCnlmClvWvLrTmtTl", "lfDDvZZSvLtDtCQZlt", "pGGhhzMRc", "mzgwGrrwhThFGPmGPcFGpw", "LjnSjLZLB", "rrfhfMRmpsghfrhGhgQr", "ZTTgvRpmfffpfgRRh", "lstcGcttdczzsW", "rjmJrmfmJM", "JbNrjGjGrMDRJ", "tftJQwCgS", "DcwvlZBmGv", "sqwDlsBvGN", "DwDrtvMHtBCvcpDcjCMFtBC", "sPLPmqhSnLQQSWq", "mVCnWCsQqqLTPWmn", "scChGddJztdNswNsdDsth", "lQLCWgWHCWbFPbbbHqLnLPHH", "QrTlTHbPnT", "fMjgFqtFWMhtjc", "sVHdHdJHwPGPGwwbpJGTbGT", "pJbHPHJGVbBswpT", "SDGSDGVPqqqQPGrTQVTQDrSr", "twCdCWMLwbtLLjBdd", "FZdbvFFb", "rzrRgqGrwgjRVq", "dTZnNpvBvnJdBpBnsJPvs", "BlplPmdlFssl", "PPHSrmfHTnmHnH", "LcsVsGtVLtLqp", "TjLJGTTJ", "tHcshJcJRhLsQscVtccJLRH", "drwdTNlwwBTpCdCdw", "wglnzCDp", "bbMjTbBvgj", "wwrcFwsrsnnCQzrCsLnRsQ", "GwLnrsmFCCzqz", "dcnwQdcdrJdSwSF", "RgmZpWlCClRlTg", "DlNVWgWpgDCCNRDJ", "sWSHsdSrHWHsbdsddBsbj", "ZNGPhRqCCRNGN", "lGlTLTGTSVVFqVTq", "QRRgRvDwWDVjm", "rHCPPHrcLF", "CldFrNFs", "qSfMgNqfp", "cbBczbtbZTcPWzc", "PBZChtzzcWPcZtBvPjtzBB", "LJPPCHtgtL", "pVGWGGjjjplhpGGVjWnl", "cjlhShjchhjGGmNVjplc"}, {"qhQLhQLMQL", "dsLtrdhrGdsq", "BfBvPzfpgfgzzWvjSzNP", "ddGldJVprrrVdNlrN", "VJvSrHqdV", "FsCPDsfBwwT", "HpHmHrhMMVrRhBnn", "GGMCpprM", "qzzJWTbqNSWZsbSTDzCW", "bZQTZgZjbZgjcQZHQPjQgZP", "TsHTjZjHFctbtHZcgZj", "DfCDmrqNfrrLSdpSfLSpNmNC", "JvTRWJTvMWRslMv", "svslRDRTMDMsVTGJGtRTs", "jhphGqBrfpBBBfZf", "BLLbQTJpBLZ", "ZMTbJQJQbf", "VtwCVCHWNqJHNNcHchCPt", "NMMDDgGLVV", "vNHghvBVGgMffvmBVDNgGV", "ZCpwZcjpjZqQCwSPp", "nnpQQDsjj", "scQsQDQQDsmmQfvfhwDmv", "GdNzRnNngbbSrrSNNzPgPGSz", "LBFrZBBLLgZwBrlfLQZtCw", "ZQZtlgNNwlrrw", "nFjRRMsPnSRcGMRMjWjWmR", "RQffbbsR", "dblFJsFSPfbDP", "mSSggpBtgn", "sszbsjZHzgdzSwzbHdsglwl", "qbdszRRRb", "HFPTmfLTGPMQrfLGmMLWWQQL", "ncZSSrtLp", "GWTScGsrZtttGTSsrzTsSZt", "pdRqMppghdRCddDg", "sVVBDVlM", "lhPlnPlChlQhJQnnzPPBCV", "NGcTTggRbdTGcbNbTTpS", "clgdBlbblBBZgJc", "cFMbWcBgclJrgd", "ppqhsqqfGjmpPhmsRGPG", "bVWPZWQbNnDV", "QQsPQQsjnbVDNVsdn", "wRRRpJMtwJwfGGrrtrwfpf", "NHCqvrsHqfff", "WWWvfvWFfFJfNqqNVFqFNfr", "GmQLCcjLjwzRLcmCG", "zBSZGNSjGGbBGZLbZZbLGB", "LssZnGCz", "TFdFcPdJJmJJc", "dThTTFJPFgg", "qJJQGhBdFJddQhFqdQBG", "MftmNzttDcVzVvt", "McrwsNbc", "qNswzsMjbhgbblbcwshlb", "RZDSPRnHSStRfnZnZTnSSTfJ", "LgFglqGCLqsGq", "cLLHsFqgvHJsLCJCGcHGsJLH", "pSSRpzjjfjpwRnSpwzRzdzjr", "mZwScwShS", "QCmmcZQWlcQcNQwlwcQCZC", "bqGMDtgDqT", "RhRQjfclMflppjGhffjZcN", "QMfMphpWjjf", "VwDCtCtVPDTzT", "SSDNwwdSddNddwdDswRDVdR", "SwRBBsBwswcwBNTDTcSR", "LthFJhFqLWWFhtq", "TdzzZLLjZzjjvddLsvv", "wbwZddTtCtv", "hFVcVHgFBWBgHgDB", "sQndTrrqdnggjqdgnTgTT", "rsgdqTjqgdHcsjHd", "mwbLWlPmzbPmlJlWbFbJ", "ggCGgtZCltDGggdCCVtZVD", "vtCGllPtDdGgtJgmlV", "jqTqHnHbjWTvH", "QdrVGBBQdVrwRQr", "MQVRwdBbBQVVbQdQsVB", "HqHnqJFWnh", "GsGBbhtsstTp", "ZrGlBTbg", "NjcwqJVrVVzQjc", "GcVWVjNNW", "jWWhjNTTcjjhTTNcqWcc", "dRrlrdwmw", "CnrWjhsTWvplg", "plCggpjjnvTpgnTTFT", "dNdbDDqDVQNR", "zdgRgWdg", "CPPRCCqQQTnnCMMCMCC", "NbJDmprrNDsDmDvJhmcvp", "BCTTFGMTCQtJtzFFFJ", "JhWQCWGGbWtWJhzTJ", "fcsrHcrpPqrcH", "zjjnMNJjrW", "rjvnqmzmnzpqjNv", "TQLLBQRRTSHTSMDCgLs", "mjvLZvjTvZTZDgTgTmQmZhZZ", "mzQjmLTLvGDvTjGbm", "pHJnnWWBBHfWHlffWfQMnM", "CCJZsRQBnszlZz", "ZBTsnsHJBJTzCQ", "FrqqVdFMqMMhMmvShFqmvVv", "mRjCdPCPCCRCQmmpRdTmTPR", "jMdQMHFMBjPMj", "wGGZNJJg", "gpzhbJszchsnsLg", "czhbzwJczcJgznNmzJNhwgsp", "HFHPVHbBbZqTZdVZBF", "WWjCllrlWGFlSFCpjgFVSWr", "WjCSFrFqCW", "wDtMMsRQTTQNRQw", "cClRCLPCgTGGLbgl", "lGbCmclmLHbbP", "vmvSNZQNrVVmZvZFvvQQv", "dFhlpwdpnpdp", "QgFBdCJBBgggCJdhgJwCggBF", "hSVjVPSPtsj", "jMMntGMh", "WWTWPCThMTFGNPTNttGwCFMh", "sjVpjHVfgs", "QFsFFbTFssNDNDsHLbTQbssP", "pbLHDspTbFsb", "ZlgClgJCJlGJzClGMtG", "ZBWZJJDWMBNVH", "VCJWfCCBMMVWrJ", "zznppPpPwlpnSS", "ZTlDcLDddddDrmrcVZDrmT", "dmcLzlrlDl", "qqRjsjLPPvp", "hVSRmzqqnn", "TmmhShSRShqSBDQSRRmQVV", "rNWWjHZNccTjZP", "JgpnrnwsTrFNLgnpNgg", "gLNpppnHsrsFgnFNp", "QzzhQWGZQVQQZDQMMzGQM", "bjSzPPzjsMsjBBMWMM", "QqbsQBMqJj", "dlhDgptthhHpVrNVdhlvhl", "scwZcBZB", "MhJMthsbcJtctZTw", "VDdPSDQnddvcpPQjdjg", "plZgSjggjSpSljSjBpSmgmWB", "pgwNpBBgHHBBBBZSjNSl", "vPvTwwrQq", "rprPjpGVpQzvp", "rsVpPrzzzHVzNllzGlVNv", "nLtfJFLFJgWl", "FthFcFrtcdhrhDQdDmdcdt", "NwdcwQhmmdmtwddBNcc", "vfvsSTVVSgVLqSsLsLzPPLV", "DzDWffDZ", "zsHrszsZssTsWrDrDVTNNH", "SjjnngqSjQQJbQjjmgmn", "dgSFCdjjWWnMSgnSjgjbj", "WWWRnjWWj", "pQLcpvpBpZpcfpvrvQvBQLv", "gCtRrttCP", "GWtgNCrtGMPP", "jqLLjLvlJ", "WgzWmvDhggzWvfmgD", "RzchcnDmvmvnRvvvcnZ", "qBqBLrBCjqBwqJwqBzLqzw", "DrplQvFQ", "FQwrZswQrwDpFpppdwfvQ", "WWSCjhWPShPz", "DrmGVgZFDm", "FFGTLffrrGf", "pwpGzjMBpvbhps", "bcvLsHVLsrsZsVvCb", "wHsHccVVLQCZZZCVH", "pzNztmjjzhhzDfNnMjnph", "lTZJttffD", "twslNTfNQntsJQZZTlN", "jgHmggmbmbblgWcGmcbcG", "GzZptzGcnGtpB", "LtZznLBBzZznLqnpqnpnznV", "CSQTQhTd", "BzJTNJHWZTWZzNF", "JNJTnJzmzzfNNTNHNJNWmnvF", "ptfwSjwtVLfwdcpCjVwwSp", "RRbRCRggRbzGCRHgRCDGB", "zcCbFHGPHPDbPHRCFGGRHHG", "LrlrmhqZhhrrLLJQhNr", "sHTGHTChhHGtHsFFbbsfHhH", "hvTRtfdfvfCbFtTFFFTtvff", "VlbDrVrWD", "DQgDgQDFtdglRtlQ", "RTtsDLDgRTsNddFT", "wSCtjqcGrrnGjC", "VwcDrjcrmVMwVwHJ", "JjBTTmcrTMrMrwmczcwwTz", "gFNsgFpHCtgCNbGZZlZCll", "fLrSrgZHDvfPDPv", "GMDHGZHSgzHPDZDSf", "wtbVwCmCthmBTGb", "QpzjSSNpBWBTQpVSSD", "zdgjVVjWg", "qcChrhlf", "wLsmGLlLMsMmMGb", "hMMPVsmlPGmG", "dbStbnntrCdSrjWnb", "CZCGPpPlHmdPblls", "CmpdGsdlDpZZZPbPCmW", "FTWVTLQMFWN", "ffzcRZznnZFf", "nFZzNZplfpfRZfBnl", "VPjVJWtJHHhQV", "HMbZTmRQTbpMdQQsR", "djmBHmZdRMdZpbjpmpjBsHdM", "ZtCzzzzvCGGccCCDCfvJrv", "wqMtJqwtvtLHvQwtLtnJMw", "LHtJBHDZDMqD", "pcPjzVdmppczrV", "djCZGZDGGgCdqZtZCZdZt", "qdGbdFDt", "JMJJlPWLHMDSDWW", "jqHBHVlHvv", "QpRqQPlBRVVBRvvtHqf", "MGcbdFTscTmgcFlgc", "WzjGBTjznBzWNzPzTzfjTzf", "GGLWGfTLTBLQMN", "ZrFsdhrtgFhttZZhts", "RhjTNhShTjHmmjjhStBmvT", "TjhBBZTSTRSbv", "snLbfffVc", "MMssTcdM", "vfdqfRsZMqssjqdd", "gwJgJGnBwFMBzD", "ZpPptplPWtppStpSlBqlq", "WtvqPTcpgPZl", "hjmfMcbDjQC", "cPCcTLcPN", "NcwjllwvQQLQlllccwclCMCT", "rgGmbSSpS", "LGsdnGbQSs", "nwdsSQQbddsSsqnqGfs", "zqRzWzZW", "QJmFmhPSmmSsQQFhsmqSFvm", "WhHQWJSqSSQmhqhm", "wtBcpdcbcBZt", "BTCRBzRDFJCBLp", "CCJJzFCMVMTzpBMMCVCMTwpT", "GvHqttftbP", "VMwMhDQDVfqPjMhwwVq", "zzjbMfqpbww", "ZcZWSSWPGGdNcFGgcR", "CqVBWVBCbHJjRdNqWq", "qMBqMdVBN", "StTfltfDftStDsQRsflDQl", "cbBdDqzND", "cPNtZcqNCqdCqPzcNZMM", "WnVSJJLLWnvJHFlFH", "VCvQVLCHHnQHWL", "qvQTLvqfGG", "ghZggcchhDhzc", "jpdWHjbRJ", "HjdbczcRpb", "SMQTCNfCTC", "blJfJGJbZlfzrRlJdbRdZld", "lbzRRhfJCJfNJhZ", "DVMFVtHFtMBFpjwM", "vmpVmRVcZjmvVvRSMZVSZ", "VMpVRgSMmZVSBZBjZVSjRBZp", "sJGtsDdGCr", "NQSDMmNlCCDMQN", "vClldmvRSnvlDlLLSlvRnQl", "cfrBtrhmtrFczzq", "JfHJgqhgJqzhJJmWfHpWpzmg", "gfcJgRHfRRpfcpRHM", "VvsWsQsW", "JCHLzCCWhtDbCC", "SCbthLtLbbQShQQLHtSST", "fJmmRfwfwMjw", "BzRhbFgfbFwzQg", "zwBvgzQRgTTTp", "WJZZWWZS", "fvlPjZFSvmvSbvfLFfFFbGZ", "GPrrmjjmjmbSjFPrF", "JtQcztMR", "ZNSqldTlVt", "ZcNlpJVZHdttZVttSZqJZd", "hvPbPmWmvvfjj", "BBqfBfwMqfBfzsh", "BZbhHJJsZjfBHZ", "FcrtQGrrVnrcFPtc", "mfwffmJDJwcfDQgfvwJDj", "wSDvmmwwgDhgf", "sFHWntbWBsBsnBHb", "SbDnbTDlDnbqS", "JcllqbTlpq", "gVhhhWgN", "mSmlQrRFG", "rTrmRrLGFTQQFmJ", "CfHMMNdDgDpMfVfVpHMqNR", "DfFfvTLwfv", "dsvrrMFvfDMTWBdFrfFF", "QgmTPtHPPJmQgQHgtqgZ"}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		got, err := Split(tt.in)
		if (err != nil) != tt.wantErr {
			t.Fatalf("error %v, wanted error %v", err, tt.wantErr)
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, wanted %v", got, tt.want)
		}
	}
}

func Test_AssignNum(t *testing.T) {
	tests := []struct {
		name    string
		param   [][]string
		want    map[string]int
		wantErr bool
	}{
		{
			name:    "First Test",
			param:   [][]string{{"vJrwpWtwJgWr", "jqHRNqRjqzjGDLGL", "PmmdzqPrV"}, {"hcsFMMfFFhFp", "rsFMfFZSrLrFZsSL", "vPwwTWBwg"}},
			want:    map[string]int{"A": 27, "B": 56, "C": 29, "D": 60, "E": 31, "F": 4096, "G": 132, "H": 68, "I": 35, "J": 144, "K": 37, "L": 608, "M": 312, "N": 80, "O": 41, "P": 336, "Q": 43, "R": 176, "S": 180, "T": 92, "U": 47, "V": 96, "W": 392, "X": 50, "Y": 51, "Z": 208, "a": 1, "b": 2, "c": 6, "d": 8, "e": 5, "f": 24, "g": 28, "h": 32, "i": 9, "j": 80, "k": 11, "l": 12, "m": 52, "n": 14, "o": 15, "p": 64, "q": 272, "r": 1152, "s": 152, "t": 40, "u": 21, "v": 88, "w": 736, "x": 24, "y": 25, "z": 104},
			wantErr: false,
		},
		{
			name:    "Real Test",
			param:   [][]string{{"jNNBMTNzvT", "VCwnVRCGHHJT", "wFJZTbRcnJCbpwpFccZC", "wDWgDfWNTvwvgFfWf", "nLnmLSnmM", "MsmsbLvtzMj", "WTqSCqWSWqSgVZqJ", "PtLGLGdd", "jFvLPPlLjvfjjffsclvP", "lLSSrfmddlNpnmLdfSPddDd", "MJMnhzBJVJwvGsMRhRh", "vzvBWGJMVGwWGGhGqVBnGzVG", "qGWLgfbWhqpLgZb", "FQwPwNPCVFQQdNcFHNcww", "wQHnwnQFPNPdjnnP", "mpVpmtPhVtP", "vzFMlMzvvR", "srqzvRRqdvzFFrDnvqldF", "CQJJLZCJLq", "RfRFFszTtzbWTFnRtFrRrW", "znsjWzstnWjbDnDbT", "jRrzzrPjL", "CTVnTlZWHTcqZBcTTTHqW", "qlWFWWHTZTlqTFWZHtBVFdlT", "cWjbcjnssMmRPWbGsWcMbf", "vqpvFHNvhvJzd", "JzqVVVVpzhVJqzhSzJhhDV", "NznNGQgS", "jjjCLwdvLcrhq", "qrrvwHqjtL", "LGPtFtTfTfTfLrrWTSWGPrr", "NhHpVpqNN", "ccBMJVDNCvvhDNpvNDpnNCVC", "dRgRhVLRl", "WJmwvvvjQvWwFBBvFJPHpwv", "FFHJjmjjjPMvfHfw", "NDTRFSnN", "JjLZWqwjHvfrvmHHvjLWWL", "mwHvHmpvwpLmjnvpvLvv", "GGhZQMsmGRfMwfp", "tFvFSTSDnDTTLt", "vzDtzMVtnLTTMTSjSDHt", "MfJtWTTMJfzB", "FvlgSHSLvmCvZjgcC", "HLhclRlHRmHvmLmclLZFSF", "scsswLQcGmQm", "bdnPSMMMClbBBPCdDDnDVMD", "bDPTlSTPPbSPjnPSb", "tFFJjMnFhdcMMJlWtdnlFc", "QRgqwwvs", "gqfvvnpfHwRmH", "fccTzVVcfSm", "wLsWfppsjBrnnwjCBZnZ", "pZrNCWLwLZLjwwW", "rZnVVjVS", "dLvQLBLFddvBLzpGmddQC", "LLCCWjGBFPmCWdmmPWLdWBpL", "ddfnQRbpldRlR", "BWBVDZWTBTDPvVTZVDBDNMbh", "NPZPDMWZMbZrBVMDDWNmBhMV", "bvPSBttgG", "NJspjJLdprzHgjrLzLNrnH", "HdRJsJrdHT", "sNVvSdTstDCtdzdzSCwTzC", "rgWlrnLLbqb", "BmrnrFPnnBLlq", "mJMqlVlttQlFVmzFQMQMbQM", "GGZgcfjrvCCZvgCZqvPH", "GZppfGHGrgGprZh", "WDHHLtRBHgDnVrWFVFB", "JlqmMpCMCGM", "cJJqSplpfSGccJGt", "tlFwbWtQFLJhlBFlWPbwm", "GCMBDDCvRpDVpRpH", "MDDvCVGzBGfpzBSNMSRR", "fzFzFHBfnvpHFbnzbHfBHZ", "sLMNSccQLMSrmLcshL", "cwLMwwNhQhsWS", "NWlqqhNNnGtNvvW", "mppCjppMrDTSgDppCDT", "pgzzTFFFrj", "ncQrhQjqjVQh", "RllWFLRf", "vFLLmfNRFSNDfS", "mmGrgwmGD", "nDsqfPCHnpntDssfJPDP", "DfPCJLpbl", "dJDNbRhNbJdhq", "LMHrttLFcMmcMFLmtG", "czZMzrzVZzHZ", "sbsJgbsm", "LLLFBzGMLjzzFtVFwww", "GllfptwjLjlGFVFjGzpFF", "SccPbfbncpcfsjbRjM", "vvvdgVhlmlwlgZLgm", "NbLNDdbLVmqqm", "HcgDBJSHTC", "VppGMwFwllwwbZZ", "wZlhtPGPLFVFlGZbVtV", "HPMMnhBHlMnMBPBHJHPWfdnB", "SrzrSScrctrwVzCSC", "CswstFRNpcwVNRrVVSVwpw", "mlFMtqjvMdqjmM", "HcNgcPLcHLwtcf", "LGGNhbNtNLDfcgcwfbgthfw", "NbvbBGNvhNhnhpbgpGfBvNg", "qqFSSqWrtSScc", "SBDSBZZZ", "qTdHSLSFRZRHHZV", "tDvflvrGttDCjlQfmCGvtCft", "tQlDCDtrfrtlGrjlQB", "MTwvsQJMvvHwVMMJMQNNJRP", "zhcmnRZmqm", "ZbmhZcdZznZfndt", "ZhhVqQTvZvVhSmQZ", "dJWHwndMHswsw", "dnwwNJnjBjwwdddnjndfW", "trlrlrZzsjRj", "QqLvvLQmLfQTLbLTfHmqHHLq", "DPDvqLqffqf", "gsnVdcBc", "RQZLRZlZDRJDCzCjblJSDjQZ", "pDlQbSlzpQ", "zgqtCtJltdGttJgVGPPJCtJv", "rRcwBnnpcmBq", "mwRqrqrmWrZnfWmBmnR", "jlzHllmPnpHlH", "bQrhDbcLgsQrtd", "QhGdGcLGwwwtcD", "qbpqvWFHbFHHsWwPqPpsVW", "MCgSSJMCzt", "QCRMRSnNtRS", "ZnZrTfsWWv", "pLGBLBgLCpgGpbdLbgMClJ", "dpJPldgMMbglFd", "CCZCQzwwdmMGDWMmhCM", "jfStHtcjqDPbPtvqv", "fcqqPPPtPSfRVBBRPRPbv", "VVlDNDgppgtNltlrJP", "LmGRmfmwCq", "cHRcCFfmfmGLJLGFcJLFwf", "vSSdnBVp", "mLHzqtGtNfGHHFNH", "GHqHHGzNDNFCfqllCFq", "CMWcwGTrvzDWzrDccDCGzTTT", "nssRttdsnhsdbFhtVFhN", "PnVPVPLFt", "tlQlqlJCCJWgF", "SDbwShHbBZZbhcBdDBDST", "BTbbScwRhwZL", "VDfPVHfDVMMfHSPSMTVfqg", "bWGGHCplWnJNpJCCnln", "HpGCnZGllzlzJWjnCbbnCnR", "fQqcfqfS", "NcCLCtPCPMtNBwdthpLhPw", "twMPBtCpLBCtwMPpJcMB", "vQwVQQVqcJvVJvCpQBCLp", "hszPRDTVD", "TtDDZsszlPllhPzmPVGssTP", "sRmGqqzzz", "cffSwNDddTdf", "GGcTwHwHv", "JSdjLJMBdMSrfwwLp", "VsRtVllstHHNllsQsHQ", "tGVGCTTTtTFVbsTTNTHsTT", "MhjTJjlS", "RHttqbNGGmbbLmLmdqbgn", "gCNmDHHRCLHn", "BvzpbBwBms", "TnNRdtlRlCd", "jltNtRCPSlPNtc", "hflJphNDmmbpfnfpl", "FgTFWqTBfFPfFqQqT", "BSSRTGWBgFggBWgBGStTG", "mDDFjjFmV", "hhRvNNCvdNMRLzhSzpp", "WCdMzdMzhMdvRvLhCRLPv", "cScMdhsDhDDdv", "NPWjJqrjJWwrSFWRJlrlNWJ", "NWNQwNPj", "fVVCVccppZMZMMC", "QrmmPhDqPhsPRhrlbgRDbbPD", "DgbPLbPsQsrbsqLbgllsQQ", "ZlsmlrZZJcQmhBhlNrsrJ", "fMFvdfVTTdjWTTTfvSjVjpj", "cdvfSWjfjpdVqSwSvMd", "sWVVmDJsNWNjcSNJZcNcZWW", "nqPRwQRgpQRPQQgMQgQLQqB", "qgrPMLLPp", "VcGjcCHcVHPrGnjQ", "bWWhfzJhvZWJzNpm", "MhJMhJbMmSZbJZ", "FCvpgDsZNsCbvvvp", "fhfWdPhhhhRQQqzdLSRHRH", "qndhdQWndftfhWStfLQQfW", "mqTCNhTNmGTLwLN", "cjsMQppsFnFslnRQQ", "stWFjcFFsJlQcQt", "fmsffcqhmqPsnTCnCc", "vHbgLRvvv", "WltLWvWl", "CtwjffWrdznRtzC", "cZFFbgPgJZDc", "SpNFJBJccNgDcTJJT", "qWzgNFqzqHNTBzFN", "rfSJRJnhhnJDjrfvRSt", "cffRvfnnVRj", "VcHhVrVCQQWh", "dDSCGSsSblwDdmLqv", "tCCwGmSqbqtws", "fvvTcWzGcCJrJGJvv", "SgFhVgllLgjLgwlwljFqVFSF", "LNqFnlDgVqllwLFLnVSgLF", "FczpzmSjVVpSQrzzcRpRcr", "TvbGGbNfGCBB", "shgdNTGgvNsfls", "JHLPLTsSllgSSPPSPLTRT", "mhFpnpFw", "prpWcvnmhmpccBB", "GcMcjDbDMM", "rLCrwNJCnwrZNLWQQwz", "zhnrnzdrCwLJCwzwC", "cnwlFrdMsggblgsrMbncwrs", "vvVVHSpQvvRQDJ", "hCmJqVvMSpqqChCJHD", "sfDNqLNpqpzCzLsDqzbCVW", "wZrlwJwFJGlFM", "wrQrgnPwM", "hDcwwGWh", "NQmLbNSZHQSHSNpb", "LNVtLSSHLSgLNr", "CRDfCbfjcnRCBVfjVMfMj", "HdTFLGsdLrzN", "JLGmmHrwJGw", "VShGpPbWj", "HqfqDfDFFJDvZRJvqZRRqHZd", "FddsFmJjj", "NRFFLtFtqF", "ljMfMBDlJHgBVHgVflf", "PlzDzjlj", "dtzZZbctPzwdlzRwlcdfRgt", "nDjHnNjjjMnfjNTM", "CDpVCfrjLnnf", "sssZSZtDfHbbdt", "jwWgmWlNQNLlcjWhgQlrQQWr", "ggLgrqljLl", "cSGBFsFcSRZSQGsgBNg", "vnCTLlrpPlH", "rJCnlmClvWvLrTmtTl", "lfDDvZZSvLtDtCQZlt", "pGGhhzMRc", "mzgwGrrwhThFGPmGPcFGpw", "LjnSjLZLB", "rrfhfMRmpsghfrhGhgQr", "ZTTgvRpmfffpfgRRh", "lstcGcttdczzsW", "rjmJrmfmJM", "JbNrjGjGrMDRJ", "tftJQwCgS", "DcwvlZBmGv", "sqwDlsBvGN", "DwDrtvMHtBCvcpDcjCMFtBC", "sPLPmqhSnLQQSWq", "mVCnWCsQqqLTPWmn", "scChGddJztdNswNsdDsth", "lQLCWgWHCWbFPbbbHqLnLPHH", "QrTlTHbPnT", "fMjgFqtFWMhtjc", "sVHdHdJHwPGPGwwbpJGTbGT", "pJbHPHJGVbBswpT", "SDGSDGVPqqqQPGrTQVTQDrSr", "twCdCWMLwbtLLjBdd", "FZdbvFFb", "rzrRgqGrwgjRVq", "dTZnNpvBvnJdBpBnsJPvs", "BlplPmdlFssl", "PPHSrmfHTnmHnH", "LcsVsGtVLtLqp", "TjLJGTTJ", "tHcshJcJRhLsQscVtccJLRH", "drwdTNlwwBTpCdCdw", "wglnzCDp", "bbMjTbBvgj", "wwrcFwsrsnnCQzrCsLnRsQ", "GwLnrsmFCCzqz", "dcnwQdcdrJdSwSF", "RgmZpWlCClRlTg", "DlNVWgWpgDCCNRDJ", "sWSHsdSrHWHsbdsddBsbj", "ZNGPhRqCCRNGN", "lGlTLTGTSVVFqVTq", "QRRgRvDwWDVjm", "rHCPPHrcLF", "CldFrNFs", "qSfMgNqfp", "cbBczbtbZTcPWzc", "PBZChtzzcWPcZtBvPjtzBB", "LJPPCHtgtL", "pVGWGGjjjplhpGGVjWnl", "cjlhShjchhjGGmNVjplc"}, {"qhQLhQLMQL", "dsLtrdhrGdsq", "BfBvPzfpgfgzzWvjSzNP", "ddGldJVprrrVdNlrN", "VJvSrHqdV", "FsCPDsfBwwT", "HpHmHrhMMVrRhBnn", "GGMCpprM", "qzzJWTbqNSWZsbSTDzCW", "bZQTZgZjbZgjcQZHQPjQgZP", "TsHTjZjHFctbtHZcgZj", "DfCDmrqNfrrLSdpSfLSpNmNC", "JvTRWJTvMWRslMv", "svslRDRTMDMsVTGJGtRTs", "jhphGqBrfpBBBfZf", "BLLbQTJpBLZ", "ZMTbJQJQbf", "VtwCVCHWNqJHNNcHchCPt", "NMMDDgGLVV", "vNHghvBVGgMffvmBVDNgGV", "ZCpwZcjpjZqQCwSPp", "nnpQQDsjj", "scQsQDQQDsmmQfvfhwDmv", "GdNzRnNngbbSrrSNNzPgPGSz", "LBFrZBBLLgZwBrlfLQZtCw", "ZQZtlgNNwlrrw", "nFjRRMsPnSRcGMRMjWjWmR", "RQffbbsR", "dblFJsFSPfbDP", "mSSggpBtgn", "sszbsjZHzgdzSwzbHdsglwl", "qbdszRRRb", "HFPTmfLTGPMQrfLGmMLWWQQL", "ncZSSrtLp", "GWTScGsrZtttGTSsrzTsSZt", "pdRqMppghdRCddDg", "sVVBDVlM", "lhPlnPlChlQhJQnnzPPBCV", "NGcTTggRbdTGcbNbTTpS", "clgdBlbblBBZgJc", "cFMbWcBgclJrgd", "ppqhsqqfGjmpPhmsRGPG", "bVWPZWQbNnDV", "QQsPQQsjnbVDNVsdn", "wRRRpJMtwJwfGGrrtrwfpf", "NHCqvrsHqfff", "WWWvfvWFfFJfNqqNVFqFNfr", "GmQLCcjLjwzRLcmCG", "zBSZGNSjGGbBGZLbZZbLGB", "LssZnGCz", "TFdFcPdJJmJJc", "dThTTFJPFgg", "qJJQGhBdFJddQhFqdQBG", "MftmNzttDcVzVvt", "McrwsNbc", "qNswzsMjbhgbblbcwshlb", "RZDSPRnHSStRfnZnZTnSSTfJ", "LgFglqGCLqsGq", "cLLHsFqgvHJsLCJCGcHGsJLH", "pSSRpzjjfjpwRnSpwzRzdzjr", "mZwScwShS", "QCmmcZQWlcQcNQwlwcQCZC", "bqGMDtgDqT", "RhRQjfclMflppjGhffjZcN", "QMfMphpWjjf", "VwDCtCtVPDTzT", "SSDNwwdSddNddwdDswRDVdR", "SwRBBsBwswcwBNTDTcSR", "LthFJhFqLWWFhtq", "TdzzZLLjZzjjvddLsvv", "wbwZddTtCtv", "hFVcVHgFBWBgHgDB", "sQndTrrqdnggjqdgnTgTT", "rsgdqTjqgdHcsjHd", "mwbLWlPmzbPmlJlWbFbJ", "ggCGgtZCltDGggdCCVtZVD", "vtCGllPtDdGgtJgmlV", "jqTqHnHbjWTvH", "QdrVGBBQdVrwRQr", "MQVRwdBbBQVVbQdQsVB", "HqHnqJFWnh", "GsGBbhtsstTp", "ZrGlBTbg", "NjcwqJVrVVzQjc", "GcVWVjNNW", "jWWhjNTTcjjhTTNcqWcc", "dRrlrdwmw", "CnrWjhsTWvplg", "plCggpjjnvTpgnTTFT", "dNdbDDqDVQNR", "zdgRgWdg", "CPPRCCqQQTnnCMMCMCC", "NbJDmprrNDsDmDvJhmcvp", "BCTTFGMTCQtJtzFFFJ", "JhWQCWGGbWtWJhzTJ", "fcsrHcrpPqrcH", "zjjnMNJjrW", "rjvnqmzmnzpqjNv", "TQLLBQRRTSHTSMDCgLs", "mjvLZvjTvZTZDgTgTmQmZhZZ", "mzQjmLTLvGDvTjGbm", "pHJnnWWBBHfWHlffWfQMnM", "CCJZsRQBnszlZz", "ZBTsnsHJBJTzCQ", "FrqqVdFMqMMhMmvShFqmvVv", "mRjCdPCPCCRCQmmpRdTmTPR", "jMdQMHFMBjPMj", "wGGZNJJg", "gpzhbJszchsnsLg", "czhbzwJczcJgznNmzJNhwgsp", "HFHPVHbBbZqTZdVZBF", "WWjCllrlWGFlSFCpjgFVSWr", "WjCSFrFqCW", "wDtMMsRQTTQNRQw", "cClRCLPCgTGGLbgl", "lGbCmclmLHbbP", "vmvSNZQNrVVmZvZFvvQQv", "dFhlpwdpnpdp", "QgFBdCJBBgggCJdhgJwCggBF", "hSVjVPSPtsj", "jMMntGMh", "WWTWPCThMTFGNPTNttGwCFMh", "sjVpjHVfgs", "QFsFFbTFssNDNDsHLbTQbssP", "pbLHDspTbFsb", "ZlgClgJCJlGJzClGMtG", "ZBWZJJDWMBNVH", "VCJWfCCBMMVWrJ", "zznppPpPwlpnSS", "ZTlDcLDddddDrmrcVZDrmT", "dmcLzlrlDl", "qqRjsjLPPvp", "hVSRmzqqnn", "TmmhShSRShqSBDQSRRmQVV", "rNWWjHZNccTjZP", "JgpnrnwsTrFNLgnpNgg", "gLNpppnHsrsFgnFNp", "QzzhQWGZQVQQZDQMMzGQM", "bjSzPPzjsMsjBBMWMM", "QqbsQBMqJj", "dlhDgptthhHpVrNVdhlvhl", "scwZcBZB", "MhJMthsbcJtctZTw", "VDdPSDQnddvcpPQjdjg", "plZgSjggjSpSljSjBpSmgmWB", "pgwNpBBgHHBBBBZSjNSl", "vPvTwwrQq", "rprPjpGVpQzvp", "rsVpPrzzzHVzNllzGlVNv", "nLtfJFLFJgWl", "FthFcFrtcdhrhDQdDmdcdt", "NwdcwQhmmdmtwddBNcc", "vfvsSTVVSgVLqSsLsLzPPLV", "DzDWffDZ", "zsHrszsZssTsWrDrDVTNNH", "SjjnngqSjQQJbQjjmgmn", "dgSFCdjjWWnMSgnSjgjbj", "WWWRnjWWj", "pQLcpvpBpZpcfpvrvQvBQLv", "gCtRrttCP", "GWtgNCrtGMPP", "jqLLjLvlJ", "WgzWmvDhggzWvfmgD", "RzchcnDmvmvnRvvvcnZ", "qBqBLrBCjqBwqJwqBzLqzw", "DrplQvFQ", "FQwrZswQrwDpFpppdwfvQ", "WWSCjhWPShPz", "DrmGVgZFDm", "FFGTLffrrGf", "pwpGzjMBpvbhps", "bcvLsHVLsrsZsVvCb", "wHsHccVVLQCZZZCVH", "pzNztmjjzhhzDfNnMjnph", "lTZJttffD", "twslNTfNQntsJQZZTlN", "jgHmggmbmbblgWcGmcbcG", "GzZptzGcnGtpB", "LtZznLBBzZznLqnpqnpnznV", "CSQTQhTd", "BzJTNJHWZTWZzNF", "JNJTnJzmzzfNNTNHNJNWmnvF", "ptfwSjwtVLfwdcpCjVwwSp", "RRbRCRggRbzGCRHgRCDGB", "zcCbFHGPHPDbPHRCFGGRHHG", "LrlrmhqZhhrrLLJQhNr", "sHTGHTChhHGtHsFFbbsfHhH", "hvTRtfdfvfCbFtTFFFTtvff", "VlbDrVrWD", "DQgDgQDFtdglRtlQ", "RTtsDLDgRTsNddFT", "wSCtjqcGrrnGjC", "VwcDrjcrmVMwVwHJ", "JjBTTmcrTMrMrwmczcwwTz", "gFNsgFpHCtgCNbGZZlZCll", "fLrSrgZHDvfPDPv", "GMDHGZHSgzHPDZDSf", "wtbVwCmCthmBTGb", "QpzjSSNpBWBTQpVSSD", "zdgjVVjWg", "qcChrhlf", "wLsmGLlLMsMmMGb", "hMMPVsmlPGmG", "dbStbnntrCdSrjWnb", "CZCGPpPlHmdPblls", "CmpdGsdlDpZZZPbPCmW", "FTWVTLQMFWN", "ffzcRZznnZFf", "nFZzNZplfpfRZfBnl", "VPjVJWtJHHhQV", "HMbZTmRQTbpMdQQsR", "djmBHmZdRMdZpbjpmpjBsHdM", "ZtCzzzzvCGGccCCDCfvJrv", "wqMtJqwtvtLHvQwtLtnJMw", "LHtJBHDZDMqD", "pcPjzVdmppczrV", "djCZGZDGGgCdqZtZCZdZt", "qdGbdFDt", "JMJJlPWLHMDSDWW", "jqHBHVlHvv", "QpRqQPlBRVVBRvvtHqf", "MGcbdFTscTmgcFlgc", "WzjGBTjznBzWNzPzTzfjTzf", "GGLWGfTLTBLQMN", "ZrFsdhrtgFhttZZhts", "RhjTNhShTjHmmjjhStBmvT", "TjhBBZTSTRSbv", "snLbfffVc", "MMssTcdM", "vfdqfRsZMqssjqdd", "gwJgJGnBwFMBzD", "ZpPptplPWtppStpSlBqlq", "WtvqPTcpgPZl", "hjmfMcbDjQC", "cPCcTLcPN", "NcwjllwvQQLQlllccwclCMCT", "rgGmbSSpS", "LGsdnGbQSs", "nwdsSQQbddsSsqnqGfs", "zqRzWzZW", "QJmFmhPSmmSsQQFhsmqSFvm", "WhHQWJSqSSQmhqhm", "wtBcpdcbcBZt", "BTCRBzRDFJCBLp", "CCJJzFCMVMTzpBMMCVCMTwpT", "GvHqttftbP", "VMwMhDQDVfqPjMhwwVq", "zzjbMfqpbww", "ZcZWSSWPGGdNcFGgcR", "CqVBWVBCbHJjRdNqWq", "qMBqMdVBN", "StTfltfDftStDsQRsflDQl", "cbBdDqzND", "cPNtZcqNCqdCqPzcNZMM", "WnVSJJLLWnvJHFlFH", "VCvQVLCHHnQHWL", "qvQTLvqfGG", "ghZggcchhDhzc", "jpdWHjbRJ", "HjdbczcRpb", "SMQTCNfCTC", "blJfJGJbZlfzrRlJdbRdZld", "lbzRRhfJCJfNJhZ", "DVMFVtHFtMBFpjwM", "vmpVmRVcZjmvVvRSMZVSZ", "VMpVRgSMmZVSBZBjZVSjRBZp", "sJGtsDdGCr", "NQSDMmNlCCDMQN", "vClldmvRSnvlDlLLSlvRnQl", "cfrBtrhmtrFczzq", "JfHJgqhgJqzhJJmWfHpWpzmg", "gfcJgRHfRRpfcpRHM", "VvsWsQsW", "JCHLzCCWhtDbCC", "SCbthLtLbbQShQQLHtSST", "fJmmRfwfwMjw", "BzRhbFgfbFwzQg", "zwBvgzQRgTTTp", "WJZZWWZS", "fvlPjZFSvmvSbvfLFfFFbGZ", "GPrrmjjmjmbSjFPrF", "JtQcztMR", "ZNSqldTlVt", "ZcNlpJVZHdttZVttSZqJZd", "hvPbPmWmvvfjj", "BBqfBfwMqfBfzsh", "BZbhHJJsZjfBHZ", "FcrtQGrrVnrcFPtc", "mfwffmJDJwcfDQgfvwJDj", "wSDvmmwwgDhgf", "sFHWntbWBsBsnBHb", "SbDnbTDlDnbqS", "JcllqbTlpq", "gVhhhWgN", "mSmlQrRFG", "rTrmRrLGFTQQFmJ", "CfHMMNdDgDpMfVfVpHMqNR", "DfFfvTLwfv", "dsvrrMFvfDMTWBdFrfFF", "QgmTPtHPPJmQgQHgtqgZ"}},
			want:    map[string]int{"A": 27, "B": 56, "C": 29, "D": 60, "E": 31, "F": 4096, "G": 132, "H": 68, "I": 35, "J": 144, "K": 37, "L": 608, "M": 312, "N": 80, "O": 41, "P": 336, "Q": 43, "R": 176, "S": 180, "T": 92, "U": 47, "V": 96, "W": 392, "X": 50, "Y": 51, "Z": 208, "a": 1, "b": 2, "c": 6, "d": 8, "e": 5, "f": 24, "g": 28, "h": 32, "i": 9, "j": 80, "k": 11, "l": 12, "m": 52, "n": 14, "o": 15, "p": 64, "q": 272, "r": 1152, "s": 152, "t": 40, "u": 21, "v": 88, "w": 736, "x": 24, "y": 25, "z": 104},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AssignNum(tt.param)
			if (err != nil) != tt.wantErr {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got: %v\n want: %v\n", got, tt.want)
			}
		})
	}
}
