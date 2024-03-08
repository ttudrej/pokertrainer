package gameobjects

type BettingRound string

// hand refers to the hand, as in one hand played being equivalent to one hand played at the table,
// as in one sequence of the 4 betting rounds, pf, f, t, r.
type hand struct {
	currentBettingRound BettingRound
}

const (
	handStart BettingRound = "Hand Start"
	pflop     BettingRound = "Pre-Flop"
	flop      BettingRound = "Flop"
	turn      BettingRound = "Turn"
	river     BettingRound = "River"
	handEnd   BettingRound = "Hand End"
)

/*
#########################################################################
#########################################################################
#########################################################################

######## ##     ## ##    ##  ######   ######
##       ##     ## ###   ## ##    ## ##    ##
##       ##     ## ####  ## ##       ##
######   ##     ## ## ## ## ##        ######
##       ##     ## ##  #### ##             ##
##       ##     ## ##   ### ##    ## ##    ##
##        #######  ##    ##  ######   ######

#########################################################################
#########################################################################
#########################################################################
*/
