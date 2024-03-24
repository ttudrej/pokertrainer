package hand_analysis

// type relativePositionPF int
// type relativePosition int // for Flop and later

// const (
// 	sb  relativePositionPF = 8 // "SB"
// 	bb  relativePositionPF = 9 // "BB"
// 	ep1 relativePositionPF = 1 // "EP1 / UTG"
// 	ep2 relativePositionPF = 2 // "EP2 / UTG + 1"
// 	ep3 relativePositionPF = 3 // "EP3 / UTG + 2"
// 	mp1 relativePositionPF = 4 // "MP1"
// 	mp2 relativePositionPF = 5 // "MP2 / HJ / Hijack"
// 	co  relativePositionPF = 6 // "LP1 / CO / Cutoff"
// 	bu  relativePositionPF = 7 // "LP2 / BU / Button"
// 	rp1 relativePosition   = 1 // SB
// 	rp2 relativePosition   = 2 // BB
// 	rp3 relativePosition   = 3 // EP1 / UTG
// 	rp4 relativePosition   = 4 // EP2 / UTG+1
// 	rp5 relativePosition   = 5 // EP3 / UTG+2
// 	rp6 relativePosition   = 6 // MP1
// 	rp7 relativePosition   = 7 // MP2 / HJ / Hijack
// 	rp8 relativePosition   = 8 // LP1 / CO / Cutoff
// 	rp9 relativePosition   = 9 // LP2 / BU / Button
// )

// // Holds record of actions a seat took during one betting cycle within a round
// type actionList [9]actionID

// // positionStatus keeps track of relevant factors related to the player's position/sequence
// // in the current hand (characteristics of this "spot"), from the point of view
// // from a specific seat in a hand.
// type positionStatusPF struct {
// 	relPos     relativePositionPF
// 	c1Actions  actionList // Actions recorded during cycle 1, ie, UTG(0), UTG1....BU, SB, BB(8)
// 	c2Actions  actionList
// 	c3Actions  actionList
// 	numBehind  int      // How many players are yet to act, at the moment the action is on us
// 	numLastBet actionID // Last bet number, 2bet, 3bet, 4bet, ...
// 	// numCalled1Bet int // How many have limped so far
// 	// numCalled2Bet int // How many called the 2 bet
// 	// numCalled3Bet int // How many called the 3 bet
// }

// // for Flop and later streets
// type positionStatus struct {
// 	relPos     relativePosition
// 	c1Actions  actionList // Actions recorded during cycle 1 on Flop or later, ie, SB(0), BB, UTG, UTG1....BU(8)
// 	c2Actions  actionList
// 	c3Actions  actionList
// 	numBehind  int      // How many players are yet to act
// 	numLastBet actionID // Last bet number, 2bet, 3bet, 4bet, ...
// }

// // #########################################################################
// /*
// ########     ##     ##                ######     ########
// ##     ##    ##     ##               ##    ##    ##     ##
// ##     ##    ##     ##               ##          ##     ##
// ########     ##     ##    #######     ######     ########
// ##     ##    ##     ##                     ##    ##     ##
// ##     ##    ##     ##               ##    ##    ##     ##
// ########      #######                 ######     ########

// ########     ########
// ##     ##    ##
// ##     ##    ##
// ########     ######
// ##           ##
// ##           ##
// ##           ##
// */

// var ps9hBUvSB_PF_bc = positionStatusPF{
// 	relPos: bu,
// 	//                     ep1, ep2, ep3, mp1, mp2, co , bu , sb , bb
// 	c1Actions:  actionList{FLD, FLD, FLD, FLD, FLD, FLD, BT2, CLL, FLD},
// 	c2Actions:  actionList{out, out, out, out, out, out, NOP, NOP, out},
// 	c3Actions:  actionList{out, out, out, out, out, out, NOP, NOP, out},
// 	numBehind:  2,
// 	numLastBet: BT2,
// }

// var ps9hBUvSB_PF_bbc = positionStatusPF{
// 	relPos: bu,
// 	//                     ep1, ep2, ep3, mp1, mp2, co , bu , sb , bb
// 	c1Actions:  actionList{FLD, FLD, FLD, FLD, FLD, FLD, BT2, BT3, FLD},
// 	c2Actions:  actionList{out, out, out, out, out, out, CLL, NOP, out},
// 	c3Actions:  actionList{out, out, out, out, out, out, NOP, NOP, out},
// 	numBehind:  2,
// 	numLastBet: BT3,
// }

// /*
// ########
// ##
// ##
// ######
// ##
// ##
// ##
// */

// var ps9hBUvSB_F_xx = positionStatus{
// 	relPos: rp9,
// 	//                     sb,  bb,  ep1, ep2, ep3, mp1, mp2, co , bu
// 	c1Actions:  actionList{CHK, out, out, out, out, out, out, out, CHK},
// 	c2Actions:  actionList{NOP, out, out, out, out, out, out, out, NOP},
// 	c3Actions:  actionList{NOP, out, out, out, out, out, out, out, NOP},
// 	numBehind:  1,
// 	numLastBet: BT0,
// }

// var ps9hBUvSB_F_xbf = positionStatus{
// 	relPos: rp9,
// 	//                     sb,  bb,  ep1, ep2, ep3, mp1, mp2, co , bu
// 	c1Actions:  actionList{CHK, out, out, out, out, out, out, out, BT1},
// 	c2Actions:  actionList{FLD, out, out, out, out, out, out, out, WIN},
// 	c3Actions:  actionList{NOP, out, out, out, out, out, out, out, NOP},
// 	numBehind:  1,
// 	numLastBet: BT1,
// }

// var ps9hBUvSB_F_xbc = positionStatus{
// 	relPos: rp9,
// 	//                     sb,  bb,  ep1, ep2, ep3, mp1, mp2, co , bu
// 	c1Actions:  actionList{CHK, out, out, out, out, out, out, out, BT1},
// 	c2Actions:  actionList{CLL, out, out, out, out, out, out, out, NOP},
// 	c3Actions:  actionList{NOP, out, out, out, out, out, out, out, NOP},
// 	numBehind:  1,
// 	numLastBet: BT1,
// }

// var ps9hBUvSB_F_bf = positionStatus{
// 	relPos: rp9,
// 	//                     sb,  bb,  ep1, ep2, ep3, mp1, mp2, co , bu
// 	c1Actions:  actionList{BT1, out, out, out, out, out, out, out, FLD},
// 	c2Actions:  actionList{WIN, out, out, out, out, out, out, out, NOP},
// 	c3Actions:  actionList{NOP, out, out, out, out, out, out, out, NOP},
// 	numBehind:  1,
// 	numLastBet: BT1,
// }

// var ps9hBUvSB_F_bc = positionStatus{
// 	relPos: rp9,
// 	//                     sb,  bb,  ep1, ep2, ep3, mp1, mp2, co , bu
// 	c1Actions:  actionList{BT1, out, out, out, out, out, out, out, CLL},
// 	c2Actions:  actionList{NOP, out, out, out, out, out, out, out, NOP},
// 	c3Actions:  actionList{NOP, out, out, out, out, out, out, out, NOP},
// 	numBehind:  1,
// 	numLastBet: BT1,
// }

// var ps9hBUvSB_F_bbf = positionStatus{
// 	relPos: rp9,
// 	//                     sb,  bb,  ep1, ep2, ep3, mp1, mp2, co , bu
// 	c1Actions:  actionList{BT1, out, out, out, out, out, out, out, BT2},
// 	c2Actions:  actionList{FLD, out, out, out, out, out, out, out, WIN},
// 	c3Actions:  actionList{NOP, out, out, out, out, out, out, out, NOP},
// 	numBehind:  1,
// 	numLastBet: BT2,
// }

// var ps9hBUvSB_F_bbc = positionStatus{
// 	relPos: rp9,
// 	//                     sb,  bb,  ep1, ep2, ep3, mp1, mp2, co , bu
// 	c1Actions:  actionList{BT1, out, out, out, out, out, out, out, BT2},
// 	c2Actions:  actionList{CLL, out, out, out, out, out, out, out, NOP},
// 	c3Actions:  actionList{NOP, out, out, out, out, out, out, out, NOP},
// 	numBehind:  1,
// 	numLastBet: BT2,
// }

// /*
// ########
//    ##
//    ##
//    ##
//    ##
//    ##
//    ##
// */

// var ps9hBUvSB_T_xx = ps9hBUvSB_F_xx
// var ps9hBUvSB_T_xbf = ps9hBUvSB_F_xbf
// var ps9hBUvSB_T_xbc = ps9hBUvSB_F_xbc
// var ps9hBUvSB_T_bf = ps9hBUvSB_F_bf
// var ps9hBUvSB_T_bc = ps9hBUvSB_F_bc
// var ps9hBUvSB_T_bbf = ps9hBUvSB_F_bbf
// var ps9hBUvSB_T_bbc = ps9hBUvSB_F_bbc

// /*
// ########
// ##     ##
// ##     ##
// ########
// ##   ##
// ##    ##
// ##     ##
// */

// var ps9hBUvSB_R_xx = ps9hBUvSB_F_xx
// var ps9hBUvSB_R_xbf = ps9hBUvSB_F_xbf
// var ps9hBUvSB_R_xbc = ps9hBUvSB_F_xbc
// var ps9hBUvSB_R_bf = ps9hBUvSB_F_bf
// var ps9hBUvSB_R_bc = ps9hBUvSB_F_bc
// var ps9hBUvSB_R_bbf = ps9hBUvSB_F_bbf
// var ps9hBUvSB_R_bbc = ps9hBUvSB_F_bbc

// // #########################################################################
// /*
// ########     ##     ##                  ########     ########
// ##     ##    ##     ##                  ##     ##    ##     ##
// ##     ##    ##     ##                  ##     ##    ##     ##
// ########     ##     ##    #######       ########     ########
// ##     ##    ##     ##                  ##     ##    ##     ##
// ##     ##    ##     ##                  ##     ##    ##     ##
// ########      #######                   ########     ########

// ########     ########
// ##     ##    ##
// ##     ##    ##
// ########     ######
// ##           ##
// ##           ##
// ##           ##
// */

// var ps9hBUvBB_PF_bc = positionStatusPF{
// 	relPos: bu,
// 	//                     ep1, ep2, ep3, mp1, mp2, co , bu , sb , bb
// 	c1Actions:  actionList{FLD, FLD, FLD, FLD, FLD, FLD, BT2, FLD, CLL},
// 	c2Actions:  actionList{out, out, out, out, out, out, NOP, out, NOP},
// 	c3Actions:  actionList{out, out, out, out, out, out, NOP, out, NOP},
// 	numBehind:  2,
// 	numLastBet: BT2,
// }

// var ps9hBUvBB_PF_bbc = positionStatusPF{
// 	relPos: bu,
// 	//                     ep1, ep2, ep3, mp1, mp2, co , bu , sb , bb
// 	c1Actions:  actionList{FLD, FLD, FLD, FLD, FLD, FLD, BT2, FLD, BT3},
// 	c2Actions:  actionList{out, out, out, out, out, out, CLL, out, NOP},
// 	c3Actions:  actionList{out, out, out, out, out, out, NOP, out, NOP},
// 	numBehind:  2,
// 	numLastBet: BT3,
// }

// /*
// ########
// ##
// ##
// ######
// ##
// ##
// ##
// */

// var ps9hBUvBB_F_xx = positionStatus{
// 	relPos: rp9,
// 	//                     sb,  bb,  ep1, ep2, ep3, mp1, mp2, co , bu
// 	c1Actions:  actionList{out, CHK, out, out, out, out, out, out, CHK},
// 	c2Actions:  actionList{out, NOP, out, out, out, out, out, out, NOP},
// 	c3Actions:  actionList{out, NOP, out, out, out, out, out, out, NOP},
// 	numBehind:  1,
// 	numLastBet: BT0,
// }

// var ps9hBUvBB_F_xbf = positionStatus{
// 	relPos: rp9,
// 	//                     sb,  bb,  ep1, ep2, ep3, mp1, mp2, co , bu
// 	c1Actions:  actionList{out, CHK, out, out, out, out, out, out, BT1},
// 	c2Actions:  actionList{out, FLD, out, out, out, out, out, out, WIN},
// 	c3Actions:  actionList{out, NOP, out, out, out, out, out, out, NOP},
// 	numBehind:  1,
// 	numLastBet: BT1,
// }

// var ps9hBUvBB_F_xbc = positionStatus{
// 	relPos: rp9,
// 	//                     sb,  bb,  ep1, ep2, ep3, mp1, mp2, co , bu
// 	c1Actions:  actionList{out, CHK, out, out, out, out, out, out, BT1},
// 	c2Actions:  actionList{out, CLL, out, out, out, out, out, out, NOP},
// 	c3Actions:  actionList{out, NOP, out, out, out, out, out, out, NOP},
// 	numBehind:  1,
// 	numLastBet: BT1,
// }

// var ps9hBUvBB_F_bf = positionStatus{
// 	relPos: rp9,
// 	//                     sb,  bb,  ep1, ep2, ep3, mp1, mp2, co , bu
// 	c1Actions:  actionList{out, BT1, out, out, out, out, out, out, FLD},
// 	c2Actions:  actionList{out, WIN, out, out, out, out, out, out, NOP},
// 	c3Actions:  actionList{out, NOP, out, out, out, out, out, out, NOP},
// 	numBehind:  1,
// 	numLastBet: BT1,
// }

// var ps9hBUvBB_F_bc = positionStatus{
// 	relPos: rp9,
// 	//                     sb,  bb,  ep1, ep2, ep3, mp1, mp2, co , bu
// 	c1Actions:  actionList{out, BT1, out, out, out, out, out, out, CLL},
// 	c2Actions:  actionList{out, NOP, out, out, out, out, out, out, NOP},
// 	c3Actions:  actionList{out, NOP, out, out, out, out, out, out, NOP},
// 	numBehind:  1,
// 	numLastBet: BT1,
// }

// var ps9hBUvBB_F_bbf = positionStatus{
// 	relPos: rp9,
// 	//                     sb,  bb,  ep1, ep2, ep3, mp1, mp2, co , bu
// 	c1Actions:  actionList{out, BT1, out, out, out, out, out, out, BT2},
// 	c2Actions:  actionList{out, FLD, out, out, out, out, out, out, WIN},
// 	c3Actions:  actionList{out, NOP, out, out, out, out, out, out, NOP},
// 	numBehind:  1,
// 	numLastBet: BT2,
// }

// var ps9hBUvBB_F_bbc = positionStatus{
// 	relPos: rp9,
// 	//                     sb,  bb,  ep1, ep2, ep3, mp1, mp2, co , bu
// 	c1Actions:  actionList{out, BT1, out, out, out, out, out, out, BT2},
// 	c2Actions:  actionList{out, CLL, out, out, out, out, out, out, NOP},
// 	c3Actions:  actionList{out, NOP, out, out, out, out, out, out, NOP},
// 	numBehind:  1,
// 	numLastBet: BT2,
// }

// /*
// ########
//    ##
//    ##
//    ##
//    ##
//    ##
//    ##
// */

// var ps9hBUvBB_T_xx = ps9hBUvBB_F_xx
// var ps9hBUvBB_T_xbf = ps9hBUvBB_F_xbf
// var ps9hBUvBB_T_xbc = ps9hBUvBB_F_xbc
// var ps9hBUvBB_T_bf = ps9hBUvBB_F_bf
// var ps9hBUvBB_T_bc = ps9hBUvBB_F_bc
// var ps9hBUvBB_T_bbf = ps9hBUvBB_F_bbf
// var ps9hBUvBB_T_bbc = ps9hBUvBB_F_bbc

// /*
// ########
// ##     ##
// ##     ##
// ########
// ##   ##
// ##    ##
// ##     ##
// */

// var ps9hBUvBB_R_xx = ps9hBUvBB_F_xx
// var ps9hBUvBB_R_xbf = ps9hBUvBB_F_xbf
// var ps9hBUvBB_R_xbc = ps9hBUvBB_F_xbc
// var ps9hBUvBB_R_bf = ps9hBUvBB_F_bf
// var ps9hBUvBB_R_bc = ps9hBUvBB_F_bc
// var ps9hBUvBB_R_bbf = ps9hBUvBB_F_bbf
// var ps9hBUvBB_R_bbc = ps9hBUvBB_F_bbc

// // sr9hBUopen0limpers     twoCardComboList //
// // sr9hBUopen1limpersEp   twoCardComboList
// // sr9hBUopen1limpersMp   twoCardComboList
// // sr9hBUopen1limpersLp   twoCardComboList
// // sr9hBUopen2limpersEpEp twoCardComboList
// // sr9hBUopen2limpersEpMp twoCardComboList
// // sr9hBUopen2limpersEpLp twoCardComboList
// // sr9hBUopen2limpersMpMp twoCardComboList
// // sr9hBUopen2limpersMpLp twoCardComboList
// // sr9hBUopen2limpersLpLP twoCardComboList

// // sr9hBBopen0limpers     twoCardComboList
// // sr9hBBopen1limpersEp   twoCardComboList // 1 limper in EP
// // sr9hBBopen1limpersMp   twoCardComboList // 1 limper in MP ...
// // sr9hBBopen1limpersLp   twoCardComboList
// // sr9hBBopen1limpersSB   twoCardComboList
// // sr9hBBopen2limpersEpEp twoCardComboList
// // sr9hBBopen2limpersEpMp twoCardComboList
// // sr9hBBopen2limpersEpLp twoCardComboList
// // sr9hBBopen2limpersEpSB twoCardComboList
// // sr9hBBopen2limpersMpMp twoCardComboList
// // sr9hBBopen2limpersMpLp twoCardComboList
// // sr9hBBopen2limpersLpLp twoCardComboList
// // sr9hBBopen2limpersLpSB twoCardComboList

// /*
// #########################################################################
// #########################################################################
// #########################################################################

// ######## ##     ## ##    ##  ######   ######
// ##       ##     ## ###   ## ##    ## ##    ##
// ##       ##     ## ####  ## ##       ##
// ######   ##     ## ## ## ## ##        ######
// ##       ##     ## ##  #### ##             ##
// ##       ##     ## ##   ### ##    ## ##    ##
// ##        #######  ##    ##  ######   ######

// #########################################################################
// #########################################################################
// #########################################################################
// */
