TestNewNetwork:
	go test -run TestNewNetwork -v

TestStart:
	go test -run TestStart -v

TestEmptyNetwork:
	go test -run TestEmptyNetwork -v

TestBroadcast:
	go test -run TestBroadcast -v

TestSingleProposer:
	go test -run TestSingleProposer -v

TestTwoProposersSameValue:
	go test -run TestTwoProposersSameValue -v

TestTwoProposersDifferentValue:
	go test -run TestTwoProposersDifferentValue -v

TestManyProposersDifferentValues:
	go test -run TestManyProposersDifferentValues -v

TestTwoAcceptors:
	go test -run TestTwoAcceptors -v

TestManyProposersManyAcceptorsSameValue:
	go test -run TestManyProposersManyAcceptorsSameValue -v

TestManyProposersManyAcceptorsDifferentValue:
	go test -run TestManyProposersManyAcceptorsDifferentValue -v

TestManyProposersManyAcceptorsSemiSameValues:
	go test -run TestManyProposersManyAcceptorsSemiSameValues -v
