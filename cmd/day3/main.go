package main

import (
	"fmt"
	"strings"
)

func split(s string) []string {
	return []string{
		s[0 : len(s)/2],
		s[len(s)/2:],
	}
}

var priority = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}

func main() {
	var elves = strings.Split(input, "\n")

	total := 0
	for i := 0; i < len(elves); i += 3 {
		total += priority[common(elves[i:i+3])]
	}
	fmt.Println(total)
}

func common(sacks []string) rune {
	var found = map[rune]map[int]bool{}
	for k := range priority {
		found[k] = map[int]bool{}
	}
	for i, s := range sacks {
		for _, r := range s {
			found[r][i] = true
			if len(found[r]) == len(sacks) {
				return r
			}
		}
	}
	return 0
}

var input = `rNZNWvMZZmDDmwqNdZrWTqhJMhhgzggBhzBJBchQzzJJ
pHlSVbVbFHgHBzzhQHqg
nVsqGpbbtDtTNmrmfZ
zrBMnbzBchshsttfbMRBgmJggmmCHGgDhDgNDGHL
VddZqQqdvSQMJHJGdCDCDDmH
pZWWllPQlPZQvZvwpSVlqlvtfswMRzBbntzRbzbfstsRzF
NnjjRlnWNSWWbGwccbcchfPfTvfjfTBBpvmdMjTfvB
FVzJtDDJDqTMlmlM
gVQZlFLlzHhLGShGww
rPZtvtFrFPgWjQvCBlcqMzlqQC
QGVDJJnLnVTCJBczqqTM
fNSSnmLDSVLhhhSNSLhGSGfVPjrFHwmQwtwWFRWRjWPHrwgt
SvmlrVrCvmNhSSVZVCrsgqPfbwGFwwwsflbbGb
QHffdnHDDQdMGbgqPwztdPds
DjBjWHfQDfTQWTBfpMBQLVmmmcCCcVhCBBBhhCmC
trLHFFQHTLHJQrflfCnLLHrRfRRPqSRPbPbbsRGqqGqhjj
mcMpNWVVNmNVsSbSJPcGhPRR
NpzNgwzZDVNZVWNpHJQLQHtQrZQHrBCl
JVCMfgJVrJtMBhhrfVVfhVsjvpFGFgjSSgFdSGGqjvjvqF
mHllHlHpmWlDSFqbdSTS
nmZRLzQnWVpctMVpQs
BrvRzWBPWbRwGRjbbRGrtrfqjCJCjCJgJsZJscFCZcJC
MnnnVMVhTMQhsccVfwqFJgqf
mMShHHppQmHrrBzwtSbWwR
pWWGJMJJwlnZSqjWmvSWZC
gtHrLttDtgFjjqRZZCrjpp
bFtbTpHFHLbfLFbHVttccttddJGQdJzTwdTzJlMnMBwwJJ
JhqHFhVMzJPQcdcVncdc
NhgfwSjwCWwltSfnrnRWZdpcPrrRnp
NNhlltBjssNBgwLFFvDmDqLzHqBB
LnFrnddfrLnMFjWzpFhcWpjpFc
ntCwgtNggCqCgCqqPPltvcjjhvmWhmvDzTzDzD
lqlVQgVCSPVllVQSNGMHHrdQsHrJJBnMHHJf
ZGZcRZNWpcHZhJfbbNblrfrgllNr
stBMtzCCsHMfFQjfSSPgtt
qmszdsCzMncdGwdWZGvH
PccqPqbhvSvvvtWNjTtWsWcscp
gRwdDzHJQgHzfdRhgHRffzwsTTjTTCjNjssCpmWWDjtCLW
zdRMwdRHhGJwgHlnGGSFvvSrnSrr
rRpMJtPwrcCTNNQNMZQm
mDWdWVddbbbmBflFhvTHjjQjfZTgZgLLfH
bhBbFFnDVhdddFBhdmpJRrzStJmwnPzcsJ
RjlpRRWzzRGRmGzlCRRlQjCgtvTJTtJrTPttrWTwhFvvVJFT
bSBdLLqbcqcLndLHZNqcZdBDPrVTDDTJSFrJJvVthTwwDS
cqVsnBfHffVdqnZccGMmCsGzQmjsjlljgz
wMzJhLtwbnMWtHcFCCFqFNNbgq
fMlMfjrRRmdmGCGVVCHcVqcVTC
MmRRRlvmQWzpvnZpwJ
gRmgMRMmRwzzmwHbwcTNqPDVBbPTZVqPNZ
fWHphpGFpfJrrhPsNTNZVsNVhT
WGfJdvltJJfHrJpRgvMRMSwRznwMmw
htJFGsGspCppCFCGthCdpmJmgmWZfqqzWzlWcfgZHgzHlg
nwVMjVcVcWlbnBlfWB
wcNDTvPPDMFJLLppDGDD
hjCBgPbvMvmQDzlWnWjm
HrHtgZRRRNwczDWwwDzsQQWW
LpTqNtFtLFqHLHRrqgFHffVVBChvhhVPBCPhbPbp
CwpbCwjGqSjVllpGCllBfhZZRDPNcPPNvLLLDSDN
WshFFWsgTHsdMzQvPczLfLZDZRcLfR
rWsJQTMhWWHdsQTgsFJgllClVpqVbqnGblCppCVr
gRBSGcBDBSJSvPQwrTFLjggQTQ
HMMnHHHZfFVFrrMT
HhlhppCNcJzCTtBT
CCffCCmRLTsQRPHQQMPF
dWdbgcDSNclbbdwdSqHsvHPQPTPJplPMFMGJ
DWbDNcqZDSWSccNTVBCzVVfmBVZnVz
BnsrrvZwBsBSJrrrqSTgJQjCbCjgbCHDJgJFjQ
hLmGlnLmGWcjGDgfFFjQdF
hhWPmhPtczWpNRmppzRhLchMsnwZvTMZvVSwwrsNwSsBvr
tDCCltNVttJhNGlMPSWdqBqSjM
RFQcpcRTpFcnFzdLmLSWjMSSBLSQ
jwzzczpFbwnHcDCsthDJJsNbst
dLRWTHSwTmTwTcTWvQNVVQCvVvNFps
GnBPtBMJBPrjGGJMjrlqChNpNlsnhVFhQsVQ
JtMtGJfrJgDJjPjRTZLdFcRZRmwSDH
VSccPJSBLgZPDLDQ
zfpLMmLsHQGqgQHnDD
zdLLMssmrdfhddcVdJtScB
VvpTVQHSqSHSHqqHJVmRJVHpgDBwDgjcDDDgZjBZBjwBZbRw
PCdssGlstdWslFPfNPrtClGjwBgBJgJNwcjBjBgZwwMBJD
tlJldhdhdsdhTqSTqVQqQq
VGqTcTqbpPwrjfbl
BvntnZNNsLZvLszSnCsvJthlfjTrZwlrjrpPlwlhfwrl
QBtNtJLvTsFdQcqWmQRR
fjcjhmjBvcvcSvcZ
HMwZtRQQpGGRgzMvLnWWnbLlSntlbv
JQPzzJHqQRqGMMQwHwzDZZhmmPfjDjmjsCZhPj
cBlZZMfBrCBMwBMCvQzTwFbQzPnbwjTbTg
WtzpVDzmtthzGFQTbTThnnTQQg
sGWstpHdpGDmdHdmGmmmJNstRMrCcBSfBSzNBNRrSRNMcMMv
mMPDVBZZLSmRdcFpjr
fggGGfbfgQStjjsdbtdt
gNqQgCQlNCCJgJHvnvnHMjPHjv
bLsRQrQsGQbLrbRZMGgbJJBJFtlFFngJphhcfBBq
jjdHCCjfVNmmmNDFcBcpBthcplFDFq
jmvvmWVjjHTCVvNjSbQGLrRzwMWsMRwfGG
sJNCsCFFCNPhCzlrSvRrvwhRjj
MMGMTwpMHGzrGczzlG
qVmwgHtDtmCdWCsNFmNJ
fmhWhjVjNpqRRJjwRw
gnGQGDDCgSsCvPlvPgnPgnPtwqbpHRHqHdJpzpQJJJRJRF
wgPGsDGPsZgGgBmBWNZNfLWWrZ
WdsCVtjWWWHRRqLLHncC
fbSpMSPSZHRRcqlpRc
cGMmJmfMPPPccZMNQPWvjTtdTjvgmdtTsggw
tPBQhHWBtQHgWQCtLwddcGnfpGpwwnbhVb
vqQzTNJJJTvRrTNFJsZrrzFlbbfcnVbbcwmGGGpVzmddcdfd
NSSqJvFFFFFQjQCjQDSDPD
rQZnVVrZmZmgSWqHrSzHPC
LGFLwcMBcllBjFNwGjltggSqSWCCzvNgSqSHtt
wdhqqGBwwqGMcDhcwdFFbbJppZbssbfZQsQsdVQm
lqBZlsjVTbVqmFrSnTFSvwncPP
zQztHfZQtWLJzPFnnQScFcFrvS
ftHJWHhfttHWffhtgLNfZDWbdqBqjbVssBDCqCdCsmClGG
MlbWFTJQFbFFzRdNjNtjdtBT
srwnrsLVHzQPQsjjSQ
gLpnwgnwnHCvcHHcvwgCvGFFhWGmFmqMMbQFQFFhlGmJ
qqNcJgJccdqhsqgsggdgqgcrtfNWNZzVbvVFzttMfzbVMZ
GLlpPpCpwPLDGvrFVWrWWbZt
DlRCDDLSjTjDjSRSjPClwnwSHHHQmmQvTJcQgvddHsqdcgmB
jmRjRbRQLLZbPnbrcTTHHHNn
MfhhmmwtvStrpnJJHc
fgqlvfhvFzMwqfvMfFWlmMvLZsdQsZVdCdLZdGQjRzdQjD
lTPcDlVdTlVVMSDfTJccVzdlmMgGBmppgBmnHGHqHqQqqQMH
ZRjWFPsLNLLrPhWNtnBBvnpGpHGpQmHnmR
CtwssCNLrsZWjrjcbfPzwJJJffDbTl
cjMvvqpJFqhShNCRQR
ldtDgQZDPdzztLZgPTtfbnStfBSbNNSbnbhhSS
TDsrzsZZZTFHmVHjcsQW
BQmQchrmBddcmZZdpSgrpswWWswVsnnnDJVnnZFnGN
TfStMPLTHvbvRVGnHGsNnJWFNV
qtvMRMMPbbPMLqRPvRTRzMjSSmprpQdBchlmmgldgjzm
nRRnvNPhrbZDLjvS
HCszMwcHHcLDrbQDWr
ptszqwdMbnnhPBqN
QbzhhfbFhBbpbzwwLjLJjSjltL
mNndGrSStHJTJLln
rDMMNVWdVpCbSbSp
tDTSTSTTTTJDwqjWqBWttdjg
nNPmVfnGfPNVLmNzfnzPVFMjdpBwWZwZHwBLBqgjqpWH
dfGPfVQGVPhGzlmnzSvsSTDJhTbTTrrSRD
ZfgtZBptBfRQNQggjjrjjwmwsQJPzrwm
TwTGGwTwzzsJzTsH
lFvwqFLhFMnqcLlVLMLfptNWppppDBDbDfbFgW
mjftBfVPjttmjcSjcPttzJlvnrwvTRrTnvwvlRrHHTHRTR
WZDWDNLFWbZbcMDWGZDbNdMCRsnTdTvdnqrHCTrvsRRvwC
DQFZLNNgtBJQcBzJ
HbZQZFVbQVpQplQZGbGchDffltfLtmdgDjggTmtm
zWzRCdnCRBRdJrzDjLhDthjLJTTtjq
CPPnwSrRdRSzCGMcZZZMwFwMZF
WBQqNQnQllwnWQlvBBMlljHTqqFdGfmTdFfcFTFFcqmP
rsRRVrZhrzbtpZRRhFDmPvfFFrfTdFHGvc
VtSCtSLbtsZVtttthCbJSWSlJlwJQggWWglvwW
QfFLWCvRfSLFCtvtFhNcqDDcGVbhGcqh
ZVgrdZZPPZZzPwdjzZhmccsqJGqDdsDDNddD
pzzwpgZzZZTznZnjZZzPVRLQLlvfSlQRSpWlCvtSQv
RtcHhRMcrHhBrrTNDVBNLqLqQqfBPm
wCbWzWbvdWCjbWppmtmNmqmLLsfsNV
lwjWdbztgHTgggnnnR
flBbzbMfbrTlrMvBCcwPggdmcdmg
VDVVRFZRZSFFhQLSGFQhjSVZCgpvPwLCzpdWWzccwdvvvwcC
hDHRGQVHHQVRZSQGbqqfNTlbHzrbbsqb
MTFdTsZpPTcMpFCPdCBmMBmRfRGBmQgQRRgt
vbDSwvhzznnbbhDWnvSzRBgQQLgLQltqtqlmwfGB
jVjhfSnNDNbzzWzjWSjrCFNpcHdpTTJddJFpsJcc
ZrrZPHfChPdDPVVdDq
vFmsbTsmSbbBJssmSBvTmmnTrnrwlWqwVlLrVTLLTWqL
JrFbpsvFBMBmzBzFStcRhjZjfCCpZNCtct
TGgRrTggwwtvtQtdCdQNqN
sJHZJVZHDBpFBZBBNzNdhzdpSzddvqhN
VZcvFsJVFvsmvssbcnrwbrnGMbMlRn
SdcdWzMJdSMWMddZJdVcmBmwrwqrrnVnVNtr
mlQHCfgbjsfQTbfCBNtVhVnntVBnVh
HLDslDDmblgHfvLHPJFSZPpDFpFFpdPS
qNqPNJvcSzGGPQnGQp
bWhbgsshZWBhltthhbWtCsZNjrzpnQnnznnjtQFrjGjVFGnn
bRDNddhNdDsZdNChmvDmmwqqvLqwSJDq
TnSfPnCSmnSgpSTmfLzfMFLWFJJLWWsBsr
jdQjcdqDVVwDcPsPzMRJMLqPqR
PGhGchjhtZlTGTHCCb
ZZRrJJqSqJwNFFphsGsLPJ
blcMCflvTTPFFNpVvsFv
CcTlltTmtmMdmCmnlllBDDSDQSwSjRDQSdswjR
MCCPNsnQFWbvvTPF
CcCVJJhjVJZRtcCclDDlbcbTcGFFDz
HpjtVwVZfpjJVhZgCVtLmrBwdMrLsNNsMmdLqB
TJTDTnrFzzdWgWGJSSMJwg
LhPVttjtLmsPqqqVsVpsjLlgWlwHvGnlHWlgHlGgwvlP
mQshLhmsnsqZcqhZqpshsLVpNTNbBfzTRBQdFRzNNFBTdbzR
ZGqMLGqvJsJsMJmd
PDVQPfPcrrcFrrzrTdgCjSSCzgszmlJjBj
PfRtVfttVcWtVJrfbGqvwqLpRRwvpppH
HmLmMSnnWnrTrnvpqFCHVGfzVFVHQj
ttsstRhhcNwbswNtdwsdNPFfjzQppQPjfGGfQVPCpR
bbsDNtDcbhstsSZLDmSSgCmnSS
tfwBBLcJVrDnqvLv
zmWWJRZhWRRRGRNdgSZGgWTvpnjvrDqvpHjjzrpnrPDnHj
NdJmSGZWRhRNsghWTJmdGfQCtllCcFMwffBftsfMQc
lTLgTghpGZJDBrnGWnnm
VlRwlHttwqmHHbDWHJ
twldzCvsRdsFFtRtSczTjSgMcfSpSzTM
pBpMBTcSlNtMcTfFCmbPDzCDLb
JgrjjJqhGZQrQrZhnJGDDCZfvPDdDzFFdzfmZL
QHhqqnrVJJPhHrnGQgwMNwMMctcWRWSBMNtNsW
FJrlhpcfDCcFWpNpwWwjNQwz
RTTvPdbjWzMbnNNM
GRZTGggGgtvjGcqrBcttcDlFhr
pMRVdVbbMMMSdWWqHpCTvTjnBBBFFGGB
smNfZgcsNrcmzggZszsgRnPGFHjBPTBTjGjPTBNj
RmwgsmgfrzzsZtfgZLQQSVWlwbdMhlwdqQ
mRRjPmLrrSmzSczSzPgVZFpTCpZCMWrZQMQrZJZT
BvdbHNdnJtvBDbqqdBlvwvqpDQMpZQFMCsQCspZTMMCZCF
nBlfbfbndJBHPfLRfmhhhhPL
ScJDFBNLLbVRqVfZ
rWrgmdMgnnBhBtnntf
CwBWWMgCwddCgwsQjsrvNvlTJzSNHwNTHFJHzS
vnddCrNpCgtjLdSdgCgCCvLnWqDhWBQhHqQHDqBhQHDHNNDl
wPTVfVTJmZGJVJGffZBwHMWlWlHlWtbQDqbl
mGsJVVJsTVTTmtJVzzTJjdSjjprzCvpSLSCjdnLg
zLNggsVHmNNsssLmwzLQZLwDRvGQBqGGDDBBvvDBDqPhRG
WrCjbtJdbFhBRglGgjqv
JWCJcWcSdWcctnJCcJJJbcbmzwwznmgLzNzmLHmHZMwsZL
JRRDNNhhszMTzNMwCG
MnHPqmgmHjPnnvjqdmjFLQwLwTLwzTwTdGLCzS
BnPPZqmcfqgqnnZmBmqjqhfWVJlRMlhWlRDlVsssbh
nmTLTqsvqnwqsvwDPnLHdNVrMMHHCBlmVdmGNV
RgRpcJhQRfQZcJbWhQpBHCjVCdjCVGdddMllHp
fczbZhzbtcZfgRRBcWSPPwFsLSDswSwTsSzw
rbFpzFCVBrrBZCjbCzHHBVdJllGDLsLrDtsswswstGJs
QNhNNnNnnQhNWSnRhnJtdpJpJtMDGsGLLtsQ
ScmRvNRNnWWvNvNvfpTccjVZbqgZgVzqHjCjTVTVVq
BTppwCwBpwwBqnjlHcLBTHnbbSbDthsSSJgsnDDRgJRD
FVGzzvrdMGSSsdtZtZgd
QvQtvtGFlBLLjLQL
gsWWsNMjwgPMPWnMjShHHZSZbmZbbmTSnb
rlCvVQrCfqffpVjQRqCCvDDTTTmmZhZTmZhThFmhhZZhqb
CDDVJpVfrJJVJLMNzMwWwLwj
nHrcsZrssPcBPtQJLJtQQCZQpV
GFWzNzNFdNbTMMqbGTqTqzqqdLCpfDQCtRVVCLtdCfQsdCCt
TlNqGTWFNmMMszhGsmFTWGFzwHnvSjgPgvgSjllBvBnvwPBB
mpMggjgMlmtjtGMwZpcSscBlcsSblhsfSdfs
zzPVDRrLrCTQNCzNRTVFNLhBhBSqdQbcfSsJBJdbjJfB
RPTRPTVNTFzVrHVDCrTHmHtwMvwWMmtwmGjWgvGv
rLMcvfHVfMgLFvfNnBBzwRbBwnrGNs
dttJjJCtdjmwzwBCRRCqcs
TddDQDJDtQJtcJFpPQHPQMvfQlFL
LQSqqpqTCSJcsDcqQMMhnnjMjppZhwHZbZ
NRtvtmgmvdBffgtVCBWVRgFbPzHbMHbnwwjMPZfHbPjzPP
RNtvCvNdgtNNmldgvCFRNVLsQLqJcQGJJrccGSlDLDLr
GdwwqqqwGVtjdPvTCplbHTPbPzPTpp
RpLmLLpFfNsgTzclhzClThgH
ZFsWZLFZJsNsnWsnRsRfnfJQGBttjdGJjBvvwjdpjjttvj
tfPzzLrrdrQlTlvn
qJRBhNhNGVRBFRTlnJvCmvmJPCCl
VVPDNchNMVFGRMFcRVBjsZZcttSLSZzzStcWtZ
pTrwTrnjtttjprTSTNTQfcjcgPsPZfPgjdgdsQ
mCmCzvzhmJDHzJDbhFCDPsgddcsfcdsbdgVRpdVs
zqJzFCDhmqvGhMmCvmGhMCGJnSlnllSBLllLMtNpWtpNBnlt
JBhJrFLhGrnJZrlcbffndnggfggf
jqmWMGGSsqCsmpjmsDQzlcHgbtdzjjlVfctjHV
GWSmSCspCsMSpRmSmqMMCBvFLJLhTTwFhRFLLBTwrv
BCdWccqcqpQqrsNgGsWMgfNW
lFttLzzLwnfsLrsNsNLG
zjNlznlwvRPZnltwvPFnZRCbmjCcqjpcpQcqVVdbdVBm
CwTbbCGNFHtHwwjSjJpzjLMdMMzT
rscqqVvWgWrZMjrlmSzzmLrM
WPqqZnPqgncnBQQVRbCDwRHGSFHPwRNw
ZQnZwWjFvdsHwBJltfmfSlsqlJ
gPprhMDTpMpPMVNqNRqNlJhltJdJ
pLGCcCrgppCrVcMpdzjvzvjLwQQzFjwzHF
NmmmvfqcvmLSQhCLvtvL
TVlWTZVJZJsFbwWbQQhtQgLFCnSgghLt
hZJTJZhwZlRJrJWHVlblMBffmqfdNMjdGdBBqqcH
GJJfLfptGqqqnsVqVVjjDnNc
mZPSvPmBCdmwdCLDshSbRnnDDhRL
gvBrBvPBPPZCTLZmwmrgQdwfTJMHGzHfWffJzFzttHWFzW
sBMvmzWzmFmNWJfffZNLfbqZbtZq
jRQVRnhhppnVhjgnDLttLqbLqLQfDLss
jRRgpGVGhwhnspgpRppwSnBvMMcWvGczGJJHdmHJmJFF
VCLHFwHMhLghHHWhFFgWNMMVzmdmbvWdJqBPJPPBppqmBdzm
SRTsjGZTsZZnSnGZGqdBmrqPvmqqqsPpmv
GvQSGtZSQllVhtLMcLLNMH
GsNdWpdVWGSHjFCWCqFFgqngvW
mRQTcrLRmZTPRLPZfqqqHbDDDgFvFnvqzQ
hfZHrwwmcZRwlLfwlmrRjMJJsVjslVNBGNjpVBBG
pllpztRqBBvvGPpG
QQhhZQbVcZQTPMWWGbvvbMHM
cwgCQCLZChQwwLZVzCrzzqNCzrDqdFPF
bgcLPvvpcbdsbpSsHRTCqsRfWfsHRm
lZlQtthrnlVMmTHqqqqHSChB
rDtlzttnlSNrMtQjZVrcgGDLLddcdcpPgPGJJd
jvGbvLLQDSGlRmmSLjlDmRQggFBrMCwWdsBFWBFjdrrWrr
PpTfcPZpNTVNpHzTzzzpPJhBcwrrhFsrMdFcMCBFhgMF
JTTqdtfzfzJpqffNdTTHGtQRnmDQGGLQQlQRbblD
CQQCshCMwgQhMdjWJFBPpbjgmmWj
SNNvcGNSZSTDtGDcczJJBmzbjBJjmppbppms
cDtfDVNTGGGNNrwLLwHdqLhfLs
ngghZCChzhNjjNbbJfdh
slPPRLlBBlVRMvRllLLHvcpcdFfJjvdFpfHfcZ
RDZPZBLmPVWDVrQtnzSTmgTwmTSg`
