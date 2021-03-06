package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

// Hexagram holds data parsed from JSON file
type Hexagram struct {
	ID    int       `json:"id"`
	Lines [6]string `json:"lines"`
	Name  string    `json:"name"`
	Desc  string    `json:"desc"`
}

// Hexagrams holds hexagrams parsed from JSON file
type Hexagrams struct {
	Hexagrams []Hexagram `json:"hexagrams"`
}

func findHexagram(a [6]string, b [6]string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func generateHexagram(coins bool) [6]string {
	var freshHexagram [6]string

	if coins {
		for i := 0; i < len(freshHexagram); i++ {
			c1 := rand.Intn(4-2) + 2
			c2 := rand.Intn(4-2) + 2
			c3 := rand.Intn(4-2) + 2
			sum := c1 + c2 + c3

			if sum == 6 {
				freshHexagram[i] = "--- X ---"
			} else if sum == 9 {
				freshHexagram[i] = "----O----"
			} else if sum == 7 {
				freshHexagram[i] = "---------"
			} else {
				freshHexagram[i] = "---   ---"
			}
		}
		return freshHexagram
	}

	marbles := [16]string{
		"--- X ---",
		"----O----", "----O----", "----O----",
		"---------", "---------", "---------", "---------", "---------",
		"---   ---", "---   ---", "---   ---", "---   ---", "---   ---",
		"---   ---", "---   ---"}

	for i := 0; i < len(freshHexagram); i++ {
		line := rand.Intn(len(marbles))
		freshHexagram[i] = marbles[line]
	}
	return freshHexagram
}

func wordWrap(text string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}

	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
	return wrapped
}

func findHxgrmManually(find string) [6]string {
	var re = regexp.MustCompile(`[xy]{6}`)
	var shape [6]string

	if !re.MatchString(find) || len(find) != 6 {
		flag.Usage()
		os.Exit(1)
	} else {
		for i := range find {
			if string(find[i]) == "x" {
				shape[i] = "---------"
			} else {
				shape[i] = "---   ---"
			}
		}
	}
	return shape
}

func printer(hexagram Hexagram, title string, quiet bool) {
	fmt.Println(title)
	for i := range hexagram.Lines {
		fmt.Printf("    %s", hexagram.Lines[len(hexagram.Lines)-1-i])
		fmt.Println()
	}
	fmt.Printf("        %v\n", hexagram.ID)
	fmt.Printf("    %v\n", hexagram.Name)
	fmt.Println()
	if !quiet {
		fmt.Println(wordWrap(hexagram.Desc, 35))
		fmt.Println()
	}
}

func isFlagPassed(flg string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == flg {
			found = true
		}
	})
	return found
}

func main() {
	var jsonData = []byte(`
	{
	    "hexagrams": [ {
	        "id":   1,
	        "lines": ["---------", "---------",  "---------",  "---------", "---------",  "---------"],
	        "name": " Force",
	        "desc": "Strength, creative energy, action; the power of heaven to create and destroy; dynamic, untiring, tenacious, enduring."
	        }, {
	        "id":   2,
	        "lines": ["---   ---", "---   ---", "---   ---", "---   ---", "---   ---", "---   ---"],
	        "name": " Field",
	        "desc": "Yield, nourish, provide; the power to give form to all things; receptive, gentle, giving, supple;\nwelcome, consent."
	        }, {
	        "id":   3,
	        "lines": ["---------", "---   ---",  "---   ---",  "---   ---", "---------", "---   ---"],
	        "name": "Sprouting",
	        "desc": "Beginning of growth and its problems; gather your strength; establish, found, assemble."
	        }, {
	        "id":   4,
	        "lines": ["---   ---",  "---------", "---   ---",  "---   ---", "---   ---",  "---------"],
	        "name": "Enveloping",
	        "desc": "Immature, young, unaware; concealed, hidden; nurture hidden growth, apprenticeship."
	        }, {
	        "id":   5,
	        "lines": ["---------", "---------", "---------", "---   ---", "---------", "---   ---"],
	        "name": "Attending",
	        "desc": "Wait for, wait on; attend to what is needed; watch for the right moment; participant in a sacrifice."
	        }, {
	        "id":   6,
	        "lines": ["---   ---",  "---------", "---   ---",  "---------", "---------", "---------"],
	        "name": "Arguing",
	        "desc": "Dispute, controversy, argument; express your position; resolve or retreat from conflict."
	        }, {
	        "id":   7,
	        "lines": ["---   ---",  "---------", "---   ---",  "---   ---", "---   ---",  "---   ---"],
	        "name": "Legions",
	        "desc": "Discipline, organize into functional units, mobilize, lead; master of arms."
	        }, {
	        "id":   8,
	        "lines": ["---   ---",  "---   ---",  "---   ---",  "---   ---", "---------", "---   ---"],
	        "name": "Grouping",
	        "desc": "Alliance, mutual support, spiritual kin; how you group things and people; changing groups."
	        }, {
	        "id":   9,
	        "lines": ["---------", "---------", "---------", "---   ---", "---------", "---------"],
	        "name": "Small Accumulating",
	        "desc": "Accumulate small things to do something great; adapt to each thing that crosses your path; nurture, tame, support, collect."
	        }, {
	        "id":   10,
	        "lines": ["---------", "---------", "---   ---",  "---------", "---------", "---------"],
	        "name": "Treading",
	        "desc": "Find and make your way, step by step; conduct, manners, salary, support."
	        }, {
	        "id":   11,
	        "lines": ["---------", "---------", "---------", "---   ---", "---   ---", "---   ---"],
	        "name": "Pervading",
	        "desc": "Prospering, expanding, great abundance and harmony; peace, communication; spring, flowering."
	        }, {
	        "id":   12,
	        "lines": ["---   ---",  "---   ---",  "---   ---",  "---------", "---------", "---------"],
	        "name": "Obstruction",
	        "desc": "Obstacle, blocked communication; decline, cut off, closed; late autumn."
	        }, {
	        "id":   13,
	        "lines": ["---------", "---   ---",  "---------", "---------", "---------", "---------"],
	        "name": "Concording People",
	        "desc": "Harmony, bring people together, share your idea or goal, welcome others, co-operate."
	        }, {
	        "id":   14,
	        "lines": ["---------", "---------", "---------", "---------", "---   ---",  "---------"],
	        "name": "Great Possessions",
	        "desc": "A powerful idea; great power to realize things; organize your efforts, concentrate; great results and achievements."
	        }, {
	        "id":   15,
	        "lines": ["---   ---", "---   ---", "---------", "---   ---", "---   ---", "---   ---"],
	        "name": "Humbling",
	        "desc": "Cut through pride and complications, keep close to fundamental things; be simple; think and speak of yourself humbly."
	        }, {
	        "id":   16,
	        "lines": ["---   ---",  "---   ---",  "---   ---",  "---------", "---   ---",  "---   ---"],
	        "name": "Providing For",
	        "desc": "Gather what you need to meet the future; able to respond immediately; enjoy, pleasure, enthusiasm, be carried away."
	        }, {
	        "id":   17,
	        "lines": ["---------", "---   ---",  "---   ---",  "---------", "---------", "---   ---"],
	        "name": "Following",
	        "desc": "Be drawn into motion; influenced by, accept guidance; move with the flow, natural and correct."
	        }, {
	        "id":   18,
	        "lines": ["---   ---",  "---------", "---------", "---   ---", "---   ---",  "---------"],
	        "name": "Corruption",
	        "desc": "Disorder, perversion or decay with roots in the past, black magic; renew, renovate, find a new beginning."
	        }, {
	        "id":   19,
	        "lines": ["---------", "---------", "---   ---",  "---   ---", "---   ---",  "---   ---"],
	        "name": "Nearing",
	        "desc": "Approach, the arrival of the new, growing; an honoured and powerful force comes nearer."
	        }, {
	        "id":   20,
	        "lines": ["---   ---",  "---   ---",  "---   ---",  "---   ---", "---------", "---------"],
	        "name": "Viewing",
	        "desc": "Look at things from a distance, contemplate, let everything come into view, divine the meaning."
	        }, {
	        "id":   21,
	        "lines": ["---------", "---   ---",  "---   ---",  "---------", "---   ---",  "---------"],
	        "name": "Gnawing And\n    Biting Through",
	        "desc": "Confront the problem, bite through the obstacle, be tenacious, reveal the essential."
	        }, {
	        "id":   22,
	        "lines": ["---------", "---   ---",  "---------", "---   ---", "---   ---",  "---------"],
	        "name": "Adorning",
	        "desc": "Make outward appearance reflect inner worth; embellish, beautify, display courage and beauty to build inner value."
	        }, {
	        "id":   23,
	        "lines": ["---   ---",  "---   ---",  "---   ---",  "---   ---", "---   ---",  "---------"],
	        "name": "Stripping",
	        "desc": "Strip away old ideas and habits, eliminate what is unusable, outmoded or worn out."
	        }, {
	        "id":   24,
	        "lines": ["---------", "---   ---",  "---   ---",  "---   ---", "---   ---",  "---   ---"],
	        "name": "Returning",
	        "desc": "Energy and spirit return after a difficult time; renewal, re-birth, re-establish; new hope."
	        }, {
	        "id":   25,
	        "lines": ["---------", "---   ---",  "---   ---",  "---------", "---------", "---------"],
	        "name": "Without Embroiling",
	        "desc": "Disentangle yourself; spontaneous, unplanned, direct; clean, pure, free from confusion or ulterior motives."
	        }, {
	        "id":   26,
	        "lines": ["---------", "---------", "---------", "---   ---", "---   ---",  "---------"],
	        "name": "Great Accumulating",
	        "desc": "Concentrate, focus on a great idea; accumulate energy, bring everything together; a time for great effort and achievement."
	        }, {
	        "id":   27,
	        "lines": ["---------", "---   ---",  "---   ---",  "---   ---", "---   ---",  "---------"],
	        "name": "  Jaws",
	        "desc": "Nourishing and being nourished, food and words; the mouth, your daily bread; take things in, swallow."
	        }, {
	        "id":   28,
	        "lines": ["---   ---",  "---------", "---------", "---------", "---------", "---   ---"],
	        "name": "Great Exceeding",
	        "desc": "A crisis; gather all your force, don't be afraid to act alone; hold on to your ideals."
	        }, {
	        "id":   29,
	        "lines": ["---   ---",  "---------", "---   ---",  "---   ---", "---------", "---   ---"],
	        "name": "Repeating The Gorge",
	        "desc": "Unavoidable danger; take the plunge, face your fear; practise, confront something repeatedly."
	        }, {
	        "id":   30,
	        "lines": ["---------", "---   ---",  "---------", "---------", "---   ---",  "---------"],
	        "name": "Radiance",
	        "desc": "Light, warmth and spreading awareness; join with, adhere to; see clearly."
	        }, {
	        "id":   31,
	        "lines": ["---   ---",  "---   ---",  "---------", "---------", "---------", "---   ---"],
	        "name": "Conjoining",
	        "desc": "Influence or stimulus to action, excite, mobilize; connection, bring together what belongs together."
	        }, {
	        "id":   32,
	        "lines": ["---   ---",  "---------", "---------", "---------", "---   ---",  "---   ---"],
	        "name": "Persevering",
	        "desc": "Continue on, endure and renew the way, constant, consistent, continue in what is right."
	        }, {
	        "id":   33,
	        "lines": ["---   ---",  "---   ---",  "---------", "---------", "---------", "---------"],
	        "name": "Retiring",
	        "desc": "Withdraw, conceal yourself, retreat; pull back in order to advance later."
	        }, {
	        "id":   34,
	        "lines": ["---------", "---------", "---------", "---------", "---   ---",  "---   ---"],
	        "name": "Great Invigorating",
	        "desc": "Great strength, the strength of the Great, have a firm purpose, focus your strength and go forward."
	        }, {
	        "id":   35,
	        "lines": ["---   ---",  "---   ---",  "---   ---",  "---------", "---   ---",  "---------"],
	        "name": "Prospering",
	        "desc": "Step into the light, advance surely, receive gifts, be promoted, spread prosperity, dawn of a new day."
	        }, {
	        "id":   36,
	        "lines": ["---------", "---   ---",  "---------", "---   ---", "---   ---",  "---   ---"],
	        "name": "Hiding Brightness",
	        "desc": "Hide your light, protect yourself, accept the difficult task."
	        }, {
	        "id":   37,
	        "lines": ["---------", "---   ---",  "---------", "---   ---", "---------", "---------"],
	        "name": "Dwelling People",
	        "desc": "Hold together, an enduring group; adapt, nourish, support; family, clan."
	        }, {
	        "id":   38,
	        "lines": ["---------", "---------", "---   ---",  "---------", "---   ---",  "---------"],
	        "name": "Diverging",
	        "desc": "Opposition, discord; change conflict into creative tension through awareness."
	        }, {
	        "id":   39,
	        "lines": ["---   ---",  "---   ---",  "---------", "---   ---", "---------", "---   ---"],
	        "name": "Difficulties",
	        "desc": "Confront obstacles; feel hampered or afflicted."
	        }, {
	        "id":   40,
	        "lines": ["---   ---",  "---------", "---   ---",  "---------", "---   ---",  "---   ---"],
	        "name": "Loosening",
	        "desc": "Solve problems, untie knots, release blocked energy; liberation, end of suffering."
	        }, {
	        "id":   41,
	        "lines": ["---------", "---------", "---   ---",  "---   ---", "---   ---",  "---------"],
	        "name": "Diminishing",
	        "desc": "Loss, decrease, sacrifice; concentrate, diminish involvements; aim at a higher goal."
	        }, {
	        "id":   42,
	        "lines": ["---------", "---   ---",  "---   ---",  "---   ---", "---------", "---------"],
	        "name": "Augmenting",
	        "desc": "Increase, expand, develop, pour in more, a fertile and expansive time."
	        }, {
	        "id":   43,
	        "lines": ["---------", "---------", "---------", "---------", "---------", "---   ---"],
	        "name": "Deciding",
	        "desc": "A critical moment, a breakthrough; decide and act clearly, clean it out and bring it to light."
	        }, {
	        "id":   44,
	        "lines": ["---   ---",  "---------", "---------", "---------", "---------", "---------"],
	        "name": "Coupling",
	        "desc": "Opening, welcoming, an intense personal encounter; meet and act through the yin, sexual intercourse."
	        }, {
	        "id":   45,
	        "lines": ["---   ---",  "---   ---",  "---   ---",  "---------", "---------", "---   ---"],
	        "name": "Clustering",
	        "desc": "Gather, assemble, collect, bunch together, crowds; a great effort brings great rewards."
	        }, {
	        "id":   46,
	        "lines": ["---   ---",  "---------", "---------", "---   ---", "---   ---",  "---   ---"],
	        "name": "Ascending",
	        "desc": "Rise to a higher level, lift yourself, advance; climb up step by step."
	        }, {
	        "id":   47,
	        "lines": ["---   ---",  "---------", "---   ---",  "---------", "---------", "---   ---"],
	        "name": "Confining",
	        "desc": "Oppression, restriction, being cut off; the moment of truth; turn inward, find a way to open communication."
	        }, {
	        "id":   48,
	        "lines": ["---   ---",  "---------", "---------", "---   ---", "---------", "---   ---"],
	        "name": "The Well",
	        "desc": "Communicate, interact, in good order; the underlying structure, network; source of life-water necessary to all."
	        }, {
	        "id":   49,
	        "lines": ["---------", "---   ---",  "---------", "---------", "---------", "---   ---"],
	        "name": "Skinning",
	        "desc": "Renew; moult, change radically, strip away the old, revolution, revolt."
	        }, {
	        "id":   50,
	        "lines": ["---   ---",  "---------", "---------", "---------", "---   ---",  "---------"],
	        "name": "The Vessel",
	        "desc": "Transformation, reach to the spiritual level; found, consecrate, imagine, contain."
	        }, {
	        "id":   51,
	        "lines": ["---------", "---   ---",  "---   ---",  "---------", "---   ---",  "---   ---"],
	        "name": " Shake",
	        "desc": "A disturbing and fertilizing shock; wake up, stir up, begin the new; return of life and love in spring."
	        }, {
	        "id":   52,
	        "lines": ["---   ---",  "---   ---",  "---------", "---   ---", "---   ---",  "---------"],
	        "name": " Bound",
	        "desc": "Calm, still, stabilize; limit or boundary, end of a cycle; become an individual."
	        }, {
	        "id":   53,
	        "lines": ["---   ---",  "---   ---",  "---------", "---   ---", "---------", "---------"],
	        "name": "Gradual Advancing",
	        "desc": "Step by step, smooth, adaptable, penetrate like water; the oldest daughter's marriage."
	        }, {
	        "id":   54,
	        "lines": ["---------", "---------", "---   ---",  "---------", "---   ---",  "---   ---"],
	        "name": "Converting The Maiden",
	        "desc": "Choice or transformation over which you have no control; realize your hidden potential; passion, desire, irregular progress; the younger daughter's marriage."
	        }, {
	        "id":   55,
	        "lines": ["---------", "---   ---",  "---------", "---------", "---   ---",  "---   ---"],
	        "name": "Abounding",
	        "desc": "Culmination, plenty, copious, profusion; generosity, opulence, full to overflowing."
	        }, {
	        "id":   56,
	        "lines": ["---   ---",  "---   ---",  "---------", "---------", "---   ---",  "---------"],
	        "name": "Sojourning",
	        "desc": "Wandering, living in exile, searching for your individual truth; outside the social net, on a quest."
	        }, {
	        "id":   57,
	        "lines": ["---   ---",  "---------", "---------", "---   ---", "---------", "---------"],
	        "name": "Gently Penetrating",
	        "desc": "Supple, flexible, subtle penetration; accept, let yourself be shaped by the situation; support or nourish from below."
	        }, {
	        "id":   58,
	        "lines": ["---------", "---------", "---   ---",  "---------", "---------", "---   ---"],
	        "name": "  Open",
	        "desc": "Communication, self-expression; pleasure, joy, interaction; persuade, exchange, the marketplace."
	        }, {
	        "id":   59,
	        "lines": ["---   ---",  "---------", "---   ---",  "---   ---", "---------", "---------"],
	        "name": "Dispersing",
	        "desc": "Dissolve, clear away, scatter, clear up; make fluid, eliminate obstacles and misundestandings."
	        }, {
	        "id":   60,
	        "lines": ["---------", "---------", "---   ---",  "---   ---", "---------", "---   ---"],
	        "name": "Articulating",
	        "desc": "Give measure, limit and form; articulate thought and speech; rhythm, interval, chapter, units."
	        }, {
	        "id":   61,
	        "lines": ["---------", "---------", "---   ---",  "---   ---", "---------", "---------"],
	        "name": "Connecting To Centre",
	        "desc": "Connection to the spirit; just, sincere, truthful; the power of a heart free of prejudice; connect the inner and outer parts of your life."
	        }, {
	        "id":   62,
	        "lines": ["---   ---",  "---   ---",  "---------", "---------", "---   ---",  "---   ---"],
	        "name": "Small Exceeding",
	        "desc": "A time of transition, adapt to each different thing; be very careful, very small; excess yin."
	        }, {
	        "id":   63,
	        "lines": ["---------", "---   ---",  "---------", "---   ---", "---------", "---   ---"],
	        "name": "Already Fording",
	        "desc": "Already underway, the action has begun; proceed actively, everything is in place and in order."
	        }, {
	        "id":   64,
	        "lines": ["---   ---",  "---------", "---   ---",  "---------", "---   ---",  "---------"],
	        "name": "Not Yet Fording",
	        "desc": "On the edge of an important change; gather your energy, everything is possible; wait for the right moment."
	        }
	    ]
	}`)

	var coins, quiet bool = false, false
	var showhex int
	var find string
	flag.BoolVar(&coins, "c", false, "Use coins method instead of marbles")
	flag.BoolVar(&quiet, "q", false, "Don't show descriptions")
	flag.IntVar(&showhex, "s", 0, "Show specific hexagram (1-64) and its description")
	flag.StringVar(&find, "f", "", "Find hexagram by its lines: x denotes Yang line, y denotes Yin line (starting from the bottom up)")

	flag.Parse()

	phex := Hexagram{}
	rhex := Hexagram{}

	var h Hexagrams

	err := json.Unmarshal(jsonData, &h)
	if err != nil {
		fmt.Println(err)
	}

	var relating = false
	var initialHxgrm, primaryShape, relatingShape [6]string
	var primaryTitle = "  Primary Figure"
	const relatingTitle string = "  Relating Figure"

	if isFlagPassed("s") {
		if showhex < 1 || showhex > 64 {
			flag.Usage()
			os.Exit(1)
		} else {
			phex.ID = h.Hexagrams[showhex-1].ID
			phex.Name = h.Hexagrams[showhex-1].Name
			phex.Lines = h.Hexagrams[showhex-1].Lines
			phex.Desc = h.Hexagrams[showhex-1].Desc
			printer(phex, "", quiet)
			os.Exit(0)
		}
	}

	rand.Seed(time.Now().UnixNano())

	if isFlagPassed("f") {
		initialHxgrm = findHxgrmManually(find)
		primaryTitle = ""
	} else {
		initialHxgrm = generateHexagram(coins)
	}

	for i := 0; i < len(initialHxgrm); i++ {
		if initialHxgrm[i] == "--- X ---" {
			primaryShape[i] = "---   ---"
			relatingShape[i] = "---------"
			relating = true
		} else if initialHxgrm[i] == "----O----" {
			primaryShape[i] = "---------"
			relatingShape[i] = "---   ---"
			relating = true
		} else if initialHxgrm[i] == "---------" {
			primaryShape[i] = "---------"
			relatingShape[i] = "---------"
		} else if initialHxgrm[i] == "---   ---" {
			primaryShape[i] = "---   ---"
			relatingShape[i] = "---   ---"
		}
	}

	for match := true; match; match = false {
		for i := 0; i < len(h.Hexagrams); i++ {
			match := findHexagram(primaryShape, h.Hexagrams[i].Lines)
			if match {
				phex.ID = h.Hexagrams[i].ID
				phex.Name = h.Hexagrams[i].Name
				if relating {
					phex.Lines = initialHxgrm
				} else {
					phex.Lines = h.Hexagrams[i].Lines
				}
				phex.Desc = h.Hexagrams[i].Desc
				break
			}
		}
	}

	if relating {
		for match := true; match; match = false {
			for i := 0; i < len(h.Hexagrams); i++ {
				match := findHexagram(relatingShape, h.Hexagrams[i].Lines)
				if match {
					rhex.ID = h.Hexagrams[i].ID
					rhex.Name = h.Hexagrams[i].Name
					rhex.Lines = h.Hexagrams[i].Lines
					rhex.Desc = h.Hexagrams[i].Desc
					break
				}
			}
		}
	}

	printer(phex, primaryTitle, quiet)
	if relating {
		printer(rhex, relatingTitle, quiet)
	}
}
