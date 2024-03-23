package hand_recording

type bettingRound string

// hand refers to the hand, as in one hand played being equivalent to one hand played at the table,
// as in one sequence of the 4 betting rounds, pf, f, t, r.
type hand struct {
	currentBettingRound bettingRound
}

const (
	handStart bettingRound = "Hand Start"
	pflop     bettingRound = "Pre-Flop"
	flop      bettingRound = "Flop"
	turn      bettingRound = "Turn"
	river     bettingRound = "River"
	handEnd   bettingRound = "Hand End"
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
