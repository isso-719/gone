package src

import (
	"embed"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"log"
	"os"
	"time"
)

func playSound(fs embed.FS) {
	f, err := fs.Open("audio/gone.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	if err := speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10)); err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

func showEmoji() {
	// アスキーアートを表示する
	ascii := "    .dL.              ..gMMMNa,        ..gMMMNa,        NNN,   uNR      NNNNNNNR\n .dHHHHHMa          .MMM#\"\"\"WM@!     .JMM#\"\"\"YMMN.     .MMMM, .MMF     .MMY\"\"\"\"\"\n MHHHHHH@7'        .MM@              dMM'     .MMb     JM#MMN..MM\\     dMM-....\n.HHHHHHHFJHHHHH[   .MM[  .MMMMMM`   .MMF      .MMF    .MMF,MMNJM#     .MMMMMMMt\n.HHHHHHHb.,        .MMN.    .MMF     MMN,    .MMM'    .MM% -MMMMF     .MM\\\n.HHHHHHHHHb         ,MMMMNMMMMD      .WMMMNMMMM@`     dMM   dMMM\\     dMMMMMMM\\\n7\"\"\"\"\"\"\"\"\"\"^           7\"\"\"\"!           ?\"\"\"\"!        77^    777      77777777\n"
	cyan := "\033[36m"
	reset := "\033[0m"
	fmt.Println(cyan + ascii + reset)
}

func noArg(fs embed.FS) {
	go showEmoji()
	playSound(fs)
}

func help() {
	fmt.Println("Gone:")
	fmt.Println(" Gone is a CLI command that sounds a \"GONE\" sound to dispel annoyances.")
	fmt.Println("Usage:")
	fmt.Println(" gone [option]")
	fmt.Println("The options are:")
	fmt.Println(" --help : Show this help.")
}

func Exec(fs embed.FS) {
	if len(os.Args) > 1 {
		help()
		return
	}
	noArg(fs)
}
