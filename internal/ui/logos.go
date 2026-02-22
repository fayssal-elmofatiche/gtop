package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Logo art adapted from onefetch (MIT License)
// https://github.com/o2sh/onefetch

type logoData struct {
	art    string
	colors []string
}

var logos = map[string]logoData{
	"Go":         goLogo,
	"Python":     pythonLogo,
	"JavaScript": javascriptLogo,
	"TypeScript": typescriptLogo,
	"Rust":       rustLogo,
	"Java":       javaLogo,
	"C":          cLogo,
	"C++":        cppLogo,
	"C#":         csharpLogo,
	"Ruby":       rubyLogo,
	"PHP":        phpLogo,
	"Swift":      swiftLogo,
	"Kotlin":     kotlinLogo,
	"Shell":      shellLogo,
	"HTML":       htmlLogo,
}

var goLogo = logoData{
	colors: []string{"#74CDDD", "#FFFFFF", "#F6D2A2"},
	art: `{0}           --==============--
{0}  .-==-.===oooo=oooooo=ooooo===--===-
{0} .==  =o={1}oGGGGGG{0}o=oo=o{1}GGGGGGG{0}G=o=  oo-
{0} -o= oo={1}G .=GGGGG{0}o=o={1}= .=GGGGG{0}=ooo o=-
{0}  .-=oo={1}o==oGGGGG{0}=oo={1}oooGGGGGo{0}=oooo.
{0}   -ooooo{1}=oooooo{0}={2}.   .{0}={1}=ooo=={0}oooooo-
{0}   -ooooooooooo{2}====_===={0}ooooooooooo=
{0}   -oooooooooooo{2}=={1}#{0}.{1}#{2}=={0}ooooooooooooo
{0}   -ooooooooooooo={1}#{0}.{1}#{0}=oooooooooooooo
{0}   .oooooooooooooooooooooooooooooooo.
{0}    oooooooooooooooooooooooooooooooo.
{2}  ..{0}oooooooooooooooooooooooooooooooo{2}..
{2}-=o-{0}=ooooooooooooooooooooooooooooooo{2}-oo.
{2}.=- {0}oooooooooooooooooooooooooooooooo{2}-.-
{0}   .oooooooooooooooooooooooooooooooo-
{0}   -oooooooooooooooooooooooooooooooo-
{0}   -oooooooooooooooooooooooooooooooo-
{0}   -oooooooooooooooooooooooooooooooo-
{0}   .oooooooooooooooooooooooooooooooo
{0}    =oooooooooooooooooooooooooooooo-
{0}    .=oooooooooooooooooooooooooooo-
{0}      -=oooooooooooooooooooooooo=.
{2}     =oo{0}====oooooooooooooooo==-{2}oo=-
{2}    .-==-    {0}.--=======---     {2}.==-`,
}

var pythonLogo = logoData{
	colors: []string{"#2F69A2", "#FFD940"},
	art: `{0}               =========
{0}            ===============
{0}           =================
{0}          ===  ==============
{0}          ===================
{0}                   ==========
{0}   ========================== {1}=======
{0} ============================ {1}========
{0}============================= {1}=========
{0}============================ {1}==========
{0}========================== {1}============
{0}============ {1}==========================
{0}========== {1}============================
{0}========= {1}=============================
{0} ======== {1}============================
{0}  ======= {1}==========================
{1}          ==========
{1}          ===================
{1}          ==============  ===
{1}           =================
{1}            ===============
{1}               =========`,
}

var javascriptLogo = logoData{
	colors: []string{"#ECE653"},
	art: `{0}JSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJS
{0}JSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJS
{0}JSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJS
{0}JSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJS
{0}JSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJS
{0}JSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJS
{0}JSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJS
{0}JSJSJSJSJSJSJSJSJ    SJSJS      JSJSJS
{0}JSJSJSJSJSJSJSJSJ    SJS          JSJS
{0}JSJSJSJSJSJSJSJSJ    SJS     JSJSJSJSJ
{0}JSJSJSJSJSJSJSJSJ    SJSJ     SJSJSJSJ
{0}JSJSJSJSJSJSJSJSJ    SJSJSJ     SJSJSJ
{0}JSJSJSJSJSJSJSJSJ    SJSJSJSJ     JSJS
{0}JSJSJSJSJSJSJSJSJ    SJSJSJSJS     JSJ
{0}JSJSJSJSJS     JS    JSJS          JSJ
{0}JSJSJSJSJSJ          SJSJSJ      SJSJS
{0}JSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJS
{0}JSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJSJS`,
}

var typescriptLogo = logoData{
	colors: []string{"#007ACC", "#FFFFFF"},
	art: `{0}TSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTS
{0}TSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTS
{0}TSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTS
{0}TSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTS
{0}TSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTS
{0}TSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTS
{0}TSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTS
{0}TSTSTSTS{1}TSTSTSTSTSTSTS{0}TSTS{1}TSTSTS{0}TSTSTS
{0}TSTSTSTS{1}TSTSTSTSTSTSTS{0}TS{1}TSTSTSTSTS{0}TSTS
{0}TSTSTSTSTSTST{1}STST{0}STSTSTS{1}TSTST{0}TSTSTSTST
{0}TSTSTSTSTSTST{1}STST{0}STSTSTST{1}STSTS{0}TSTSTSTS
{0}TSTSTSTSTSTST{1}STST{0}STSTSTSTST{1}STSTS{0}TSTSTS
{0}TSTSTSTSTSTST{1}STST{0}STSTSTSTSTST{1}STSTS{0}TSTS
{0}TSTSTSTSTSTST{1}STST{0}STSTSTSTSTSTS{1}TSTST{0}TST
{0}TSTSTSTSTSTST{1}STST{0}STSTSTST{1}STSTSTSTST{0}STS
{0}TSTSTSTSTSTST{1}STST{0}STSTSTSSTS{1}TSTSTS{0}TSTST
{0}TSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTS
{0}TSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTSTS`,
}

var rustLogo = logoData{
	colors: []string{"#E43717", "#FFFFFF"},
	art: `{0}                 R RR RR
{0}              R RRRRRRRR R          R
{0} R RR       R RRRRRRRRRRRRR R      RR
{0}rR RRR    R RRRRRRRRRRRRRRRRR R   RRR R
{0}RRR RR   RRRRRRRRRRRRRRRRRRRRRRR  RRRRR
{0} RRRRR  RRRRRRRRRRRRRRRRRRRRRRRR  RRRR
{0}  RRR RRRRRRRRRRRRRRRRRRRRRRRRRRRR RR
{0}    R  RRRRRRRRRR{1}=  {0}RR{1} = {0}RRRRRRRRRRR
{0}     RRRRRRRRRRRR{1}=  {0}RR{1} = {0}RRRRRRRRRR
{0}      RRRRRRRRRRR   RR   RRRRRRRRRR
{0}     RR==RRRRRRRRRRRRRRRRRRRRRR===RR
{0}     RR =  ==RRRRRRR  RRRRRR==  = RR
{0}      RR =     ===========     = RR
{0}       RR                        R
{0}        R                       R
{0}         R`,
}

var javaLogo = logoData{
	colors: []string{"#F44336", "#1665C0"},
	art: `{0}                  |
{0}                 ||
{0}               |||
{0}             ||||    ||
{0}           ||||| ||||
{0}          ||||  |||
{0}         ||||  |||
{0}         |||    |||
{0}          |||    |||
{0}            ||    ||
{0}              |   |
{1}   ####               #    ##
{1}    ################       ##
{1}       #                   ##
{1}      ################   ###
{1}
{1}       ##############
{1}####      #######          #
{1}#####                   ####
{1}   #####################      #
{1}                          ###
{1}          ###############`,
}

var cLogo = logoData{
	colors: []string{"#649AD2", "#004283", "#00599D", "#FFFFFF"},
	art: `{0}                 ++++++
{0}              ++++++++++++
{0}          ++++++++++++++++++++
{0}       ++++++++++++++++++++++++++
{0}    ++++++++++++++++++++++++++++++++
{0} +++++++++++++{3}************{0}+++++++++++++
{0}+++++++++++{3}******************{0}++++++++{2};;;
{0}+++++++++{3}**********************{0}++{2};;;;;;;
{0}++++++++{3}*********{0}++++++{3}******{2};;;;;;;;;;;
{0}+++++++{3}********{0}++++++++++{3}**{2};;;;;;;;;;;;;
{0}+++++++{3}*******{0}+++++++++{2};;;;;;;;;;;;;;;;;
{0}+++++++{3}******{0}+++++++{2};;;;;;;;;;;;;;;;;;;;
{0}+++++++{3}*******{0}+++{1}:::::{2};;;;;;;;;;;;;;;;;;
{0}+++++++{3}********{1}::::::::::{3}**{2};;;;;;;;;;;;;
{0}++++++++{3}*********{1}::::::{3}******{2};;;;;;;;;;;
{0}++++++{1}:::{3}**********************{1}::{2};;;;;;;
{0}+++{1}::::::::{3}******************{1}::::::::{2};;;
{1} :::::::::::::{3}************{1}:::::::::::::
{1}    ::::::::::::::::::::::::::::::::
{1}       ::::::::::::::::::::::::::
{1}          ::::::::::::::::::::
{1}              ::::::::::::
{1}                 ::::::`,
}

var cppLogo = logoData{
	colors: []string{"#649AD2", "#004283", "#00599D", "#FFFFFF"},
	art: `{0}                 ++++++
{0}              ++++++++++++
{0}          ++++++++++++++++++++
{0}       ++++++++++++++++++++++++++
{0}    ++++++++++++++++++++++++++++++++
{0} +++++++++++++{3}************{0}+++++++++++++
{0}+++++++++++{3}******************{0}++++++++{2};;;
{0}+++++++++{3}**********************{0}++{2};;;;;;;
{0}++++++++{3}*********{0}++++++{3}******{2};;;;;;;;;;;
{0}+++++++{3}********{0}++++++++++{3}**{2};;;;;;;;;;;;;
{0}+++++++{3}*******{0}+++++++++{2};;;;;;{3}**{2};;;;{3}**{2};;;
{0}+++++++{3}******{0}+++++++{2};;;;;;;;{3}****{2};;{3}****{2};;
{0}+++++++{3}*******{0}+++{1}:::::{2};;;;;;;{3}**{2};;;;{3}**{2};;;
{0}+++++++{3}********{1}::::::::::{3}**{2};;;;;;;;;;;;;
{0}++++++++{3}*********{1}::::::{3}******{2};;;;;;;;;;;
{0}++++++{1}:::{3}**********************{1}::{2};;;;;;;
{0}+++{1}::::::::{3}******************{1}::::::::{2};;;
{1} :::::::::::::{3}************{1}:::::::::::::
{1}    ::::::::::::::::::::::::::::::::
{1}       ::::::::::::::::::::::::::
{1}          ::::::::::::::::::::
{1}              ::::::::::::
{1}                 ::::::`,
}

var csharpLogo = logoData{
	colors: []string{"#9B4F97", "#67217A", "#803788", "#FFFFFF"},
	art: `{0}                 ++++++
{0}              ++++++++++++
{0}          ++++++++++++++++++++
{0}       ++++++++++++++++++++++++++
{0}    ++++++++++++++++++++++++++++++++
{0} +++++++++++++{3}************{0}+++++++++++++
{0}+++++++++++{3}******************{0}++++++++{2};;;
{0}+++++++++{3}**********************{0}++{2};;;;;;;
{0}++++++++{3}*********{0}++++++{3}******{2};;;;;;;;;;;
{0}+++++++{3}********{0}++++++++++{3}**{2};;;{3}**{2};;;{3}**{2};;;
{0}+++++++{3}*******{0}+++++++++{2};;;;;;{3}*********{2};;
{0}+++++++{3}******{0}+++++++{2};;;;;;;;;;{3}**{2};;;{3}**{2};;;
{0}+++++++{3}*******{0}+++{1}:::::{2};;;;;;;{3}*********{2};;
{0}+++++++{3}********{1}::::::::::{3}**{2};;;{3}**{2};;;{3}**{2};;;
{0}++++++++{3}*********{1}::::::{3}******{2};;;;;;;;;;;
{0}++++++{1}:::{3}**********************{1}::{2};;;;;;;
{0}+++{1}::::::::{3}******************{1}::::::::{2};;;
{1} :::::::::::::{3}************{1}:::::::::::::
{1}    ::::::::::::::::::::::::::::::::
{1}       ::::::::::::::::::::::::::
{1}          ::::::::::::::::::::
{1}              ::::::::::::
{1}                 ::::::`,
}

var rubyLogo = logoData{
	colors: []string{"#F30301"},
	art: `{0}                -+sdmhmMMhyhNMMddm'
{0}            .smdy+:  '-smmNy/    :h\
{0}          -yNs.          mmhmo.    Md
{0}        -hNo'           oM- .oNh:  :s
{0}      :dm+'            .Ns    '/dmo.d
{0}     +Mo'              hN        yMMM
{0}   's+N/              /dMmsssoo++///NM
{0}  'hN-             +md/sM.        hNM
{0} .dm.            +mh:  .Ns      .dm+M
{0} +My          .oNy-     sM.    :Nd.+M
{0} +MM+     .:/sNd-       .Ns   sNo  sM
{0}'+MyMyhdddysmMoydmhs+:smdsM.+Nh-   hN
{0}'+M.MM:'   -Mo           hMNm/     md
{0}'+M  hm.   dm'         '+Nmdm+     My
{0}'/M:  md' /M/        -sNy:  -yNs. .Mo
{0} 'Nh   Nh'Nh      :smd+.      .sNyoM/
{0}   'h:mh:MmNss:+sdNds:-::///++oosyN-`,
}

var phpLogo = logoData{
	colors: []string{"#777BB3", "#FFFFFF"},
	art: `{0}            ################
{0}      ##########{1}/  |{0}##############
{0}   #############{1}|  |{0}#################
{0} #####{1}/   __   \|   __   \/   __   \{0}###
{0}######{1}|  |{0}##{1}|  ||  |{0}##{1}|  ||  |{0}##{1}|  |{0}####
{0}######{1}|  |{0}##{1}/  ||  |{0}##{1}|  ||  |{0}##{1}/  |{0}####
{0} #####{1}|   ____ /|__|{0}##{1}|__||   ____ /{0}###
{0}   ###{1}|  |{0}################{1}|  |{0}#######
{1}      |_ /{0}################{1}|_ /{0}####
{0}            ################`,
}

var swiftLogo = logoData{
	colors: []string{"#F88134", "#F97732", "#F96D30", "#FA632E", "#FA592C", "#FB502A", "#FB4628", "#FC3C26", "#FC3224", "#FD2822"},
	art: `{0}                         :
{0}                          ::
{1}                           :::
{1}          :                ::::
{2}     :     :                ::::
{2}      :     ::              :::::
{3}       ::    :::             :::::
{3}        :::    :::           ::::::
{4}          :::   :::          :::::::
{4}           ::::  ::::        :::::::
{5}            :::::::::::      ::::::::
{5}              :::::::::::   :::::::::
{5}               ::::::::::::::::::::::
{6}                :::::::::::::::::::::
{6}                  :::::::::::::::::::
{6}:                   :::::::::::::::::
{7} ::                   ::::::::::::::
{7}   ::::              ::::::::::::::::
{7}    ::::::::::::::::::::::::::::::::::
{8}      :::::::::::::::::::::::::::::::::
{8}        :::::::::::::::::::::::::::::::
{8}          ::::::::::::::::::::::   :::::
{9}             .::::::::::::::.         ::`,
}

var kotlinLogo = logoData{
	colors: []string{"#7F52FF", "#E44857", "#C711E1"},
	art: `{0}KOTLIN{2}KOTLINKOTLINKO{1}TLINKOTLINKOTLINKOTL
{0}KOTLINKO{2}TLINKOTLIN{1}KOTLINKOTLINKOTLINKO
{0}KOTLINKOTL{2}INKOTL{1}INKOTLINKOTLINKOTLIN
{0}KOTLINKOTLIN{2}KO{1}TLINKOTLINKOTLINKOTL
{0}KOTLINKOTLIN{1}KOTLINKOTLINKOTLINKO
{0}KOTLINKOTL{1}INKOTLINKOTLINKOTLIN
{0}KOTLINKO{1}TLINKOTLINKOTLINKOTL
{0}KOTLIN{1}KOTLINKOTLINKOTLINKO
{0}KOTL{1}INKOTLINKOTLINKOTLIN
{0}KO{1}TLINKOTLINKOTLINKOTL
{1}KOTLINKOTLINKOTLINKO{2}TL
{2}KO{1}TLINKOTLINKOTLIN{2}KOTLIN
{2}KOTL{1}INKOTLINKOTL{2}INKOTLINKO
{2}KOTLIN{1}KOTLINKO{2}TLINKOTLINKOTL
{2}KOTLINKO{1}TLIN{0}K{2}OTLINKOTLINKOTLIN
{2}KOTLINKOTL{0}INKOT{2}LINKOTLINKOTLINKO
{2}KOTLINKO{0}TLINKOTLI{2}NKOTLINKOTLINKOTL
{2}KOTLIN{0}KOTLINKOTLINK{2}OTLINKOTLINKOTLIN
{2}KOTL{0}INKOTLINKOTLINKOT{2}LINKOTLINKOTLINKO
{2}KO{0}TLINKOTLINKOTLINKOTLI{2}NKOTLINKOTLINKOTL`,
}

var shellLogo = logoData{
	colors: []string{"#FFFFFF", "#89E051"},
	art: `{0}             _._
{0}         _.-'   '-._
{0}     _.-'           '-._
{0} _.-'                   '-._
{0}|                        _,-|
{0}|                    _,-'+++|
{0}|                _,-'+++++++|
{0}|             ,-'+++++++++++|
{0}|             |++++ ++++++++|
{0}|             |+++   +++++++|
{0}|             |++  +++++++++|
{0}|             |++++  +++{1}**{0}++|
{0}|             |++   ++{1}**{0}++++|
{0}'-,_          |+++ ++++++_,-'
{0}    '-,_      |++++++_,-'
{0}        '-,_  |++_,-'
{0}            '-|-'`,
}

var htmlLogo = logoData{
	colors: []string{"#E34C26", "#FFFFFF"},
	art: `{1}  ##  ##  ######  ##   ##  ##
{1}  ##  ##    ##    ### ###  ##
{1}  ######    ##    ## # ##  ##
{1}  ##  ##    ##    ##   ##  ##
{1}  ##  ##    ##    ##   ##  ######
{0}(((((((((((((((((((((((((((((((((((
{0}(((((((((((((((((/////////////(((((
{0}(((((((((((((((((/////////////(((((
{0}(((((((                     //(((((
{0} ((((((                     //((((
{0} ((((((    ((((((/////////////((((
{0} ((((((     (((((/////////////((((
{0} ((((((                    ///((((
{0}  (((((                    ///(((
{0}  (((((((((((((((//////    ///(((
{0}  ((((((    (((((/////     ///(((
{0}  ((((((                   ///(((
{0}   (((((((               /////((
{0}   ((((((((((((((/////////////((
{0}   ((((((((((((((//////(((((((((
{0}          (((((((((((((((`,
}

var defaultLogo = logoData{
	colors: []string{"#F0883E"},
	art: `{0}   _____ _____ ___  ____
{0}  / ____|_   _/ _ \|  _ \
{0} | |  _   | || | | | |_) |
{0} | |_| |  | || |_| |  __/
{0}  \____|  |_| \___/|_|`,
}

func getLanguageLogo(language string) logoData {
	if logo, ok := logos[language]; ok {
		return logo
	}
	return defaultLogo
}

func renderColoredArt(art string, colors []string) string {
	if len(colors) == 0 {
		return art
	}

	var result strings.Builder
	lines := strings.Split(art, "\n")
	for i, line := range lines {
		if i > 0 {
			result.WriteString("\n")
		}
		currentColor := 0
		var segment strings.Builder
		j := 0
		for j < len(line) {
			if j+2 < len(line) && line[j] == '{' && line[j+1] >= '0' && line[j+1] <= '9' && line[j+2] == '}' {
				// Flush current segment
				if segment.Len() > 0 {
					style := lipgloss.NewStyle().Foreground(lipgloss.Color(colors[currentColor]))
					result.WriteString(style.Render(segment.String()))
					segment.Reset()
				}
				newColor := int(line[j+1] - '0')
				if newColor < len(colors) {
					currentColor = newColor
				}
				j += 3
				continue
			}
			segment.WriteByte(line[j])
			j++
		}
		// Flush remaining segment
		if segment.Len() > 0 {
			style := lipgloss.NewStyle().Foreground(lipgloss.Color(colors[currentColor]))
			result.WriteString(style.Render(segment.String()))
		}
	}
	return result.String()
}
