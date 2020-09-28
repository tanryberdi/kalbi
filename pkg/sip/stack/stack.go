package stack

//import "fmt"
import "github.com/marv2097/siprocket"
import "Kalbi/pkg/transport"
import "Kalbi/pkg/log"

//SipStack has multiple protocol listning points
type SipStack struct {
	ListeningPoints []transport.ListeningPoint
	OutputPoints    []chan siprocket.SipMsg
	Alive           bool
}

//CreateListenPoint creates listening point to the event dispatcher
func (ed *SipStack) CreateListenPoint(protocol string, host string, port int) transport.ListeningPoint {
	listenpoint := transport.NewTransportListenPoint(protocol, host, port)
	ed.ListeningPoints = append(ed.ListeningPoints, listenpoint)
    return listenpoint
}

//AddChannel give the ability to add channels for callbacks on each request
func (ed *SipStack) AddChannel(c chan siprocket.SipMsg) {
	ed.OutputPoints = append(ed.OutputPoints, c)
}

func (ed *SipStack) IsAlive() bool {
    return ed.Alive
}

//Start starts the sip stack
func (ed *SipStack) Start() {
	log.Log.Info("Starting Sip Stack")
    ed.Alive = true
	for ed.Alive == true {
		for _, listeningPoint := range ed.ListeningPoints {
			msg := listeningPoint.Read()
			for _, OutputPoint := range ed.OutputPoints {
				OutputPoint <- *msg
			}

		}
	}

}
