package data

import (
	"fmt"
	"math/rand"

	"example.com/types"
)

var namen []string

func Main() {
	initArray()
	//fmt.Print("\nEs folgen alle verfügbaren Namen:")
	//fmt.Print("\n", namen)
	//fmt.Print("\nEnde Der Liste\n")

}

func Hello(name string) string {
	// Return a greeting that embeds the name in a messge.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func GiveHero() types.Hero {
	h := types.Hero{}
	h.Id = rand.Intn(len(namen) + 1)
	h.Name = namen[rand.Intn(len(namen)+1)]
	return h
}

func initArray() {
	namen = append(namen, "	Ben	")
	namen = append(namen, "	Luca	")
	namen = append(namen, "	Paul	")
	namen = append(namen, "	Jonas	")
	namen = append(namen, "	Finn	")
	namen = append(namen, "	Leon	")
	namen = append(namen, "	Luis	")
	namen = append(namen, "	Lukas	")
	namen = append(namen, "	Maximilian	")
	namen = append(namen, "	Felix	")
	namen = append(namen, "	Noah	")
	namen = append(namen, "	Elias	")
	namen = append(namen, "	Julian	")
	namen = append(namen, "	Max	")
	namen = append(namen, "	Tim	")
	namen = append(namen, "	Moritz	")
	namen = append(namen, "	Henry	")
	namen = append(namen, "	Niklas	")
	namen = append(namen, "	Philipp	")
	namen = append(namen, "	Jakob	")
	namen = append(namen, "	Tom	")
	namen = append(namen, "	Krause	")
	namen = append(namen, "	Emil	")
	namen = append(namen, "	Alexander	")
	namen = append(namen, "	David	")
	namen = append(namen, "	Oskar	")
	namen = append(namen, "	Fabian	")
	namen = append(namen, "	Anton	")
	namen = append(namen, "	Erik	")
	namen = append(namen, "	Raphael	")
	namen = append(namen, "	Matteo	")
	namen = append(namen, "	Leo	")
	namen = append(namen, "	Mats	")
	namen = append(namen, "	Simon	")
	namen = append(namen, "	Jannik	")
	namen = append(namen, "	Lennard	")
	namen = append(namen, "	Liam	")
	namen = append(namen, "	Linus	")
	namen = append(namen, "	Hannes	")
	namen = append(namen, "	Mika	")
	namen = append(namen, "	Vincent	")
	namen = append(namen, "	Adrian	")
	namen = append(namen, "	Jonathan	")
	namen = append(namen, "	Theo	")
	namen = append(namen, "	Nico	")
	namen = append(namen, "	Till	")
	namen = append(namen, "	Benjamin	")
	namen = append(namen, "	Florian	")
	namen = append(namen, "	Marlon	")
	namen = append(namen, "	Julius	")
	namen = append(namen, "	Nick	")
	namen = append(namen, "	Samuel	")
	namen = append(namen, "	Nils	")
	namen = append(namen, "	Johannes	")
	namen = append(namen, "	Jona	")
	namen = append(namen, "	Carl	")
	namen = append(namen, "	Daniel	")
	namen = append(namen, "	Lennox	")
	namen = append(namen, "	Aaron	")
	namen = append(namen, "	Mattis	")
	namen = append(namen, "	Ole	")
	namen = append(namen, "	Leonard	")
	namen = append(namen, "	Constantin	")
	namen = append(namen, "	Sebastian	")
	namen = append(namen, "	Jannis	")
	namen = append(namen, "	Colin	")
	namen = append(namen, "	Joel	")
	namen = append(namen, "	Tobias	")
	namen = append(namen, "	Lenny	")
	namen = append(namen, "	Milan	")
	namen = append(namen, "	Johann	")
	namen = append(namen, "	Joshua	")
	namen = append(namen, "	Dominic	")
	namen = append(namen, "	Maxim	")
	namen = append(namen, "	John	")
	namen = append(namen, "	Mohammed	")
	namen = append(namen, "	Timo	")
	namen = append(namen, "	Robin	")
	namen = append(namen, "	Valentin	")
	namen = append(namen, "	Jayden	")
	namen = append(namen, "	Benedikt	")
	namen = append(namen, "	Justus	")
	namen = append(namen, "	Levin	")
	namen = append(namen, "	Damian	")
	namen = append(namen, "	Phil	")
	namen = append(namen, "	Toni	")
	namen = append(namen, "	Levi	")
	namen = append(namen, "	Jamie	")
	namen = append(namen, "	Lian	")
	namen = append(namen, "	Kilian	")
	namen = append(namen, "	Malte	")
	namen = append(namen, "	Noel	")
	namen = append(namen, "	Jason	")
	namen = append(namen, "	Bennet	")
	namen = append(namen, "	Tyler	")
	namen = append(namen, "	Gabriel	")
	namen = append(namen, "	Sam	")
	namen = append(namen, "	Michael	")
	namen = append(namen, "	Artur	")
	namen = append(namen, "	Bastian	")
	namen = append(namen, "	Bruno	")
	namen = append(namen, "	Lasse	")
	namen = append(namen, "	Marc	")
	namen = append(namen, "	Pepe	")
	namen = append(namen, "	Luke	")
	namen = append(namen, "	Oliver	")
	namen = append(namen, "	Marvin	")
	namen = append(namen, "	Emilio	")
	namen = append(namen, "	Silas	")
	namen = append(namen, "	Emir	")
	namen = append(namen, "	Lars	")
	namen = append(namen, "	Leopold	")
	namen = append(namen, "	Richard	")
	namen = append(namen, "	Lias	")
	namen = append(namen, "	Elia	")
	namen = append(namen, "	Matti	")
	namen = append(namen, "	Theodor	")
	namen = append(namen, "	Christian	")
	namen = append(namen, "	Jannes	")
	namen = append(namen, "	Tristan	")
	namen = append(namen, "	Leandro	")
	namen = append(namen, "	Marcel	")
	namen = append(namen, "	Frederik	")
	namen = append(namen, "	Connor	")
	namen = append(namen, "	Manuel	")
	namen = append(namen, "	Dean	")
	namen = append(namen, "	Adam	")
	namen = append(namen, "	Diego	")
	namen = append(namen, "	Piet	")
	namen = append(namen, "	Franz	")
	namen = append(namen, "	Fritz	")
	namen = append(namen, "	Hugo	")
	namen = append(namen, "	Michel	")
	namen = append(namen, "	Ilyas	")
	namen = append(namen, "	Nicolas	")
	namen = append(namen, "	Matthias	")
	namen = append(namen, "	Dennis	")
	namen = append(namen, "	Jeremy	")
	namen = append(namen, "	Neo	")
	namen = append(namen, "	Finnley	")
	namen = append(namen, "	Marco	")
	namen = append(namen, "	Pascal	")
	namen = append(namen, "	Fabio	")
	namen = append(namen, "	Ludwig	")
	namen = append(namen, "	Malik	")
	namen = append(namen, "	Nikita	")
	namen = append(namen, "	Henrik	")
	namen = append(namen, "	Lionel	")
	namen = append(namen, "	Martin	")
	namen = append(namen, "	Clemens	")
	namen = append(namen, "	Lorenz	")
	namen = append(namen, "	Milo	")
	namen = append(namen, "	Ian	")
	namen = append(namen, "	Lenn	")
	namen = append(namen, "	Arne	")
	namen = append(namen, "	Emilian	")
	namen = append(namen, "	Yusuf	")
	namen = append(namen, "	Lio	")
	namen = append(namen, "	Jasper	")
	namen = append(namen, "	Lion	")
	namen = append(namen, "	Maik	")
	namen = append(namen, "	Ferdinand	")
	namen = append(namen, "	Thomas	")
	namen = append(namen, "	Bela	")
	namen = append(namen, "	Konrad	")
	namen = append(namen, "	Marius	")
	namen = append(namen, "	Hendrik	")
	namen = append(namen, "	Julien	")
	namen = append(namen, "	Eddi	")
	namen = append(namen, "	Friedrich	")
	namen = append(namen, "	Can	")
	namen = append(namen, "	Ali	")
	namen = append(namen, "	Berat	")
	namen = append(namen, "	Maurice	")
	namen = append(namen, "	Andreas	")
	namen = append(namen, "	Taylor	")
	namen = append(namen, "	Kevin	")
	namen = append(namen, "	Alessio	")
	namen = append(namen, "	Kai	")
	namen = append(namen, "	Laurens	")
	namen = append(namen, "	Patrick	")
	namen = append(namen, "	Laurin	")
	namen = append(namen, "	Janne	")
	namen = append(namen, "	Justin	")
	namen = append(namen, "	Markus	")
	namen = append(namen, "	Carlo	")
	namen = append(namen, "	Kian	")
	namen = append(namen, "	Leonardo	")
	namen = append(namen, "	Willi	")
	namen = append(namen, "	Benno	")
	namen = append(namen, "	Devin	")
	namen = append(namen, "	Luan	")
	namen = append(namen, "	Gustav	")
	namen = append(namen, "	Leonhard	")
	namen = append(namen, "	Mert	")
	namen = append(namen, "	Chris	")
	namen = append(namen, "	Thore	")
	namen = append(namen, "	Leander	")
	namen = append(namen, "	Magnus	")
	namen = append(namen, "	Robert	")
	namen = append(namen, "	Nevio	")
	namen = append(namen, "	Ryan	")
	namen = append(namen, "	Yasin	")
	namen = append(namen, "	Fiete	")
	namen = append(namen, "	Henning	")
	namen = append(namen, "	Arian	")
	namen = append(namen, "	Roman	")
	namen = append(namen, "	Korbinian	")
	namen = append(namen, "	Carlos	")
	namen = append(namen, "	Jonte	")
	namen = append(namen, "	Alessandro	")
	namen = append(namen, "	Peter	")
	namen = append(namen, "	Deniz	")
	namen = append(namen, "	Nino	")
	namen = append(namen, "	Alex	")
	namen = append(namen, "	Antonio	")
	namen = append(namen, "	Mailo	")
	namen = append(namen, "	Brian	")
	namen = append(namen, "	Amir	")
	namen = append(namen, "	Christopher	")
	namen = append(namen, "	Thilo	")
	namen = append(namen, "	Charlie	")
	namen = append(namen, "	Damien	")
	namen = append(namen, "	Mehmet	")
	namen = append(namen, "	Ricardo	")
	namen = append(namen, "	Curt	")
	namen = append(namen, "	Dario	")
	namen = append(namen, "	Joris	")
	namen = append(namen, "	Victor	")
	namen = append(namen, "	Darian	")
	namen = append(namen, "	Josef	")
	namen = append(namen, "	Christoph	")
	namen = append(namen, "	Georg	")
	namen = append(namen, "	Jack	")
	namen = append(namen, "	Kaan	")
	namen = append(namen, "	Mick	")
	namen = append(namen, "	Enno	")
	namen = append(namen, "	Kjell	")
	namen = append(namen, "	Aiden	")
	namen = append(namen, "	Domenic	")
	namen = append(namen, "	Jesper	")
	namen = append(namen, "	Enes	")
	namen = append(namen, "	Ömer	")
	namen = append(namen, "	Titus	")
	namen = append(namen, "	Hamza	")
	namen = append(namen, "	Mustafa	")
	namen = append(namen, "	Mikail	")
	namen = append(namen, "	Torben	")
	namen = append(namen, "	Jaron	")
	namen = append(namen, "	Ruben	")
	namen = append(namen, "	Bjarne	")
	namen = append(namen, "	Miran	")
	namen = append(namen, "	Stefan	")
	namen = append(namen, "	Claas	")
	namen = append(namen, "	Sascha	")
	namen = append(namen, "	Tammo	")
	namen = append(namen, "	Cedric	")
	namen = append(namen, "	Edgar	")
	namen = append(namen, "	Jerome	")
	namen = append(namen, "	Leif	")
	namen = append(namen, "	Leonas	")
	namen = append(namen, "	Lino	")
	namen = append(namen, "	Romeo	")
	namen = append(namen, "	Andre	")
	namen = append(namen, "	Nathan	")
	namen = append(namen, "	Tino	")
	namen = append(namen, "	William	")
	namen = append(namen, "	Hanno	")
	namen = append(namen, "	Sami	")
	namen = append(namen, "	Ahmet	")
	namen = append(namen, "	Miguel	")
	namen = append(namen, "	Steven	")
	namen = append(namen, "	Emin	")
	namen = append(namen, "	Lean	")
	namen = append(namen, "	Mirac	")
	namen = append(namen, "	Semih	")
	namen = append(namen, "	Sinan	")
	namen = append(namen, "	Etienne	")
	namen = append(namen, "	Ibrahim	")
	namen = append(namen, "	Mario	")
	namen = append(namen, "	Timon	")
	namen = append(namen, "	Xaver	")
	namen = append(namen, "	Armin	")
	namen = append(namen, "	Efe	")
	namen = append(namen, "	Janosch	")
	namen = append(namen, "	Kerem	")
	namen = append(namen, "	Mio	")
	namen = append(namen, "	Wilhelm	")
	namen = append(namen, "	Albert	")
	namen = append(namen, "	Erwin	")
	namen = append(namen, "	Hans	")
	namen = append(namen, "	Marian	")
	namen = append(namen, "	Anthony	")
	namen = append(namen, "	Cem	")
	namen = append(namen, "	Emre	")
	namen = append(namen, "	Eymen	")
	namen = append(namen, "	Leonidas	")
	namen = append(namen, "	Aras	")
	namen = append(namen, "	Ensar	")
	namen = append(namen, "	Kenan	")
	namen = append(namen, "	Kuzey	")
	namen = append(namen, "	Lutz	")
	namen = append(namen, "	Selim	")
	namen = append(namen, "	Tamme	")
	namen = append(namen, "	Valentino	")
	namen = append(namen, "	Danny	")
	namen = append(namen, "	Emanuel	")
	namen = append(namen, "	Giuliano	")
	namen = append(namen, "	Hassan	")
	namen = append(namen, "	Kerim	")
	namen = append(namen, "	Umut	")
	namen = append(namen, "	Amin	")
	namen = append(namen, "	Arda	")
	namen = append(namen, "	Danilo	")
	namen = append(namen, "	Eren	")
	namen = append(namen, "	Mattes	")
	namen = append(namen, "	Vince	")
	namen = append(namen, "	Arvid	")
	namen = append(namen, "	Darius	")
	namen = append(namen, "	Dustin	")
	namen = append(namen, "	Jake	")
	namen = append(namen, "	Jarne	")
	namen = append(namen, "	Jim	")
	namen = append(namen, "	Marten	")
	namen = append(namen, "	Sean	")
	namen = append(namen, "	James	")
	namen = append(namen, "	Jean	")
	namen = append(namen, "	Lucien	")
	namen = append(namen, "	Rayan	")
	namen = append(namen, "	Elian	")
	namen = append(namen, "	Emirhan	")
	namen = append(namen, "	Furkan	")
	namen = append(namen, "	Jonne	")
	namen = append(namen, "	Kalle	")
	namen = append(namen, "	Karim	")
	namen = append(namen, "	Milian	")
	namen = append(namen, "	Timur	")
	namen = append(namen, "	Damon	")
	namen = append(namen, "	Enrico	")
	namen = append(namen, "	Marek	")
	namen = append(namen, "	Quentin	")
	namen = append(namen, "	Alwin	")
	namen = append(namen, "	Angelo	")
	namen = append(namen, "	Jesse	")
	namen = append(namen, "	Otto	")
	namen = append(namen, "	Samir	")
	namen = append(namen, "	Yassin	")
	namen = append(namen, "	Bilal	")
	namen = append(namen, "	Caspar	")
	namen = append(namen, "	Jannek	")
	namen = append(namen, "	Jarno	")
	namen = append(namen, "	Maddox	")
	namen = append(namen, "	Mahir	")
	namen = append(namen, "	Marlo	")
	namen = append(namen, "	Rico	")
	namen = append(namen, "	Tjark	")
	namen = append(namen, "	Elija	")
	namen = append(namen, "	Iven	")
	namen = append(namen, "	Joscha	")
	namen = append(namen, "	Nikolai	")
	namen = append(namen, "	Rocco	")
	namen = append(namen, "	Sven	")
	namen = append(namen, "	Berkay	")
	namen = append(namen, "	Dion	")
	namen = append(namen, "	Gregor	")
	namen = append(namen, "	Jano	")
	namen = append(namen, "	Koray	")
	namen = append(namen, "	Ramon	")
	namen = append(namen, "	Sandro	")
	namen = append(namen, "	Taylan	")
	namen = append(namen, "	Davin	")
	namen = append(namen, "	Francesco	")
	namen = append(namen, "	Jamal	")
	namen = append(namen, "	Jordan	")
	namen = append(namen, "	Lewis	")
	namen = append(namen, "	Omar	")
	namen = append(namen, "	Pius	")
	namen = append(namen, "	Taha	")
	namen = append(namen, "	Veit	")
	namen = append(namen, "	Amar	")
	namen = append(namen, "	Eduard	")
	namen = append(namen, "	Gianluca	")
	namen = append(namen, "	Ismail	")
	namen = append(namen, "	Joost	")
	namen = append(namen, "	Lucian	")
	namen = append(namen, "	Miko	")
	namen = append(namen, "	Sirac	")
	namen = append(namen, "	Thies	")
	namen = append(namen, "	Alfred	")
	namen = append(namen, "	Dylan	")
	namen = append(namen, "	Eray	")
	namen = append(namen, "	Flynn	")
	namen = append(namen, "	Hauke	")
	namen = append(namen, "	Logan	")
	namen = append(namen, "	Melvin	")
	namen = append(namen, "	Younes	")
	namen = append(namen, "	Ilja	")
	namen = append(namen, "	Jon	")
	namen = append(namen, "	Mete	")
	namen = append(namen, "	Tamino	")
	namen = append(namen, "	Alan	")
	namen = append(namen, "	Arjen	")
	namen = append(namen, "	Jenke	")
	namen = append(namen, "	Johnny	")
	namen = append(namen, "	Keno	")
	namen = append(namen, "	Loris	")
	namen = append(namen, "	Milow	")
	namen = append(namen, "	Santino	")
	namen = append(namen, "	Tiago	")
	namen = append(namen, "	Burak	")
	namen = append(namen, "	Fabrizio	")
	namen = append(namen, "	Gian	")
	namen = append(namen, "	Hennes	")
	namen = append(namen, "	Kirill	")
	namen = append(namen, "	Lorik	")
	namen = append(namen, "	Luiz	")
	namen = append(namen, "	Peer	")
	namen = append(namen, "	Talha	")
	namen = append(namen, "	Tizian	")
	namen = append(namen, "	Tommy	")
	namen = append(namen, "	Yunus	")
	namen = append(namen, "	Aidan	")
	namen = append(namen, "	Baran	")
	namen = append(namen, "	Björn	")
	namen = append(namen, "	Cornelius	")
	namen = append(namen, "	Dorian	")
	namen = append(namen, "	Hagen	")
	namen = append(namen, "	Leano	")
	namen = append(namen, "	Quinn	")
	namen = append(namen, "	Youssef	")
	namen = append(namen, "	Benny	")
	namen = append(namen, "	Bent	")
	namen = append(namen, "	Ege	")
	namen = append(namen, "	Gero	")
	namen = append(namen, "	Ivan	")
	namen = append(namen, "	Kimi	")
	namen = append(namen, "	Levent	")
	namen = append(namen, "	Luuk	")
	namen = append(namen, "	Maris	")
	namen = append(namen, "	Miro	")
	namen = append(namen, "	Sören	")
	namen = append(namen, "	Stanley	")
	namen = append(namen, "	Vito	")
	namen = append(namen, "	Yigit	")
	namen = append(namen, "	Batuhan	")
	namen = append(namen, "	Edwin	")
	namen = append(namen, "	Jakub	")
	namen = append(namen, "	Juri	")
	namen = append(namen, "	Kiyan	")
	namen = append(namen, "	Meo	")
	namen = append(namen, "	Musa	")
	namen = append(namen, "	Raul	")
	namen = append(namen, "	Ron	")
	namen = append(namen, "	Rüzgar	")
	namen = append(namen, "	Said	")
	namen = append(namen, "	Sammy	")
	namen = append(namen, "	Santiago	")
	namen = append(namen, "	Severin	")
	namen = append(namen, "	Adem	")
	namen = append(namen, "	Adriano	")
	namen = append(namen, "	Alexandros	")
	namen = append(namen, "	Ansgar	")
	namen = append(namen, "	August	")
	namen = append(namen, "	Bo	")
	namen = append(namen, "	Dejan	")
	namen = append(namen, "	Eyüp	")
	namen = append(namen, "	Fridolin	")
	namen = append(namen, "	Hüseyin	")
	namen = append(namen, "	Ilay	")
	namen = append(namen, "	Jay	")
	namen = append(namen, "	Jimmy	")
	namen = append(namen, "	Josua	")
	namen = append(namen, "	Merlin	")
	namen = append(namen, "	Nathanael	")
	namen = append(namen, "	Quirin	")
	namen = append(namen, "	Azad	")
	namen = append(namen, "	Davide	")
	namen = append(namen, "	Denny	")
	namen = append(namen, "	Evan	")
	namen = append(namen, "	Giuseppe	")
	namen = append(namen, "	Jari	")
	namen = append(namen, "	Jascha	")
	namen = append(namen, "	Neven	")
	namen = append(namen, "	Pierre	")
	namen = append(namen, "	Rick	")
	namen = append(namen, "	Tarik	")
	namen = append(namen, "	Vinzenz	")
	namen = append(namen, "	Ahmad	")
	namen = append(namen, "	Amon	")
	namen = append(namen, "	Andrej	")
	namen = append(namen, "	Aurel	")
	namen = append(namen, "	Bosse	")
	namen = append(namen, "	Demian	")
	namen = append(namen, "	Ethan	")
	namen = append(namen, "	Joe	")
	namen = append(namen, "	Joey	")
	namen = append(namen, "	Nelio	")
	namen = append(namen, "	René	")
	namen = append(namen, "	Salvatore	")
	namen = append(namen, "	Simeon	")
	namen = append(namen, "	Hannah	")
	namen = append(namen, "	Sofia	")
	namen = append(namen, "	Anna	")
	namen = append(namen, "	Lea	")
	namen = append(namen, "	Emilia	")
	namen = append(namen, "	Marie	")
	namen = append(namen, "	Lena	")
	namen = append(namen, "	Leonie	")
	namen = append(namen, "	Emily	")
	namen = append(namen, "	Lina	")
	namen = append(namen, "	Amelie	")
	namen = append(namen, "	Sophie	")
	namen = append(namen, "	Lilly	")
	namen = append(namen, "	Luisa	")
	namen = append(namen, "	Johanna	")
	namen = append(namen, "	Laura	")
	namen = append(namen, "	Nele	")
	namen = append(namen, "	Lara	")
	namen = append(namen, "	Maja	")
	namen = append(namen, "	Charlotte	")
	namen = append(namen, "	Clara	")
	namen = append(namen, "	Leni	")
	namen = append(namen, "	Sarah	")
	namen = append(namen, "	Pia	")
	namen = append(namen, "	Mila	")
	namen = append(namen, "	Alina	")
	namen = append(namen, "	Lisa	")
	namen = append(namen, "	Lotta	")
	namen = append(namen, "	Ida	")
	namen = append(namen, "	Julia	")
	namen = append(namen, "	Greta	")
	namen = append(namen, "	Mathilda	")
	namen = append(namen, "	Melina	")
	namen = append(namen, "	Zoe	")
	namen = append(namen, "	Frieda	")
	namen = append(namen, "	Lia	")
	namen = append(namen, "	Paula	")
	namen = append(namen, "	Marlene	")
	namen = append(namen, "	Ella	")
	namen = append(namen, "	Emely	")
	namen = append(namen, "	Jana	")
	namen = append(namen, "	Victoria	")
	namen = append(namen, "	Josephine	")
	namen = append(namen, "	Finja	")
	namen = append(namen, "	Isabell	")
	namen = append(namen, "	Helena	")
	namen = append(namen, "	Isabella	")
	namen = append(namen, "	Elisa	")
	namen = append(namen, "	Amy	")
	namen = append(namen, "	Mara	")
	namen = append(namen, "	Mira	")
	namen = append(namen, "	Katharina	")
	namen = append(namen, "	Jasmin	")
	namen = append(namen, "	Stella	")
	namen = append(namen, "	Lucy	")
	namen = append(namen, "	Luise	")
	namen = append(namen, "	Antonia	")
	namen = append(namen, "	Annika	")
	namen = append(namen, "	Fiona	")
	namen = append(namen, "	Pauline	")
	namen = append(namen, "	Nora	")
	namen = append(namen, "	Eva	")
	namen = append(namen, "	Jule	")
	namen = append(namen, "	Magdalena	")
	namen = append(namen, "	Luna	")
	namen = append(namen, "	Merle	")
	namen = append(namen, "	Carla	")
	namen = append(namen, "	Maria	")
	namen = append(namen, "	Nina	")
	namen = append(namen, "	Theresa	")
	namen = append(namen, "	Melissa	")
	namen = append(namen, "	Franziska	")
	namen = append(namen, "	Martha	")
	namen = append(namen, "	Milena	")
	namen = append(namen, "	Chiara	")
	namen = append(namen, "	Ronja	")
	namen = append(namen, "	Carlotta	")
	namen = append(namen, "	Elena	")
	namen = append(namen, "	Romy	")
	namen = append(namen, "	Mina	")
	namen = append(namen, "	Helene	")
	namen = append(namen, "	Selina	")
	namen = append(namen, "	Annabell	")
	namen = append(namen, "	Paulina	")
	namen = append(namen, "	Vanessa	")
	namen = append(namen, "	Maila	")
	namen = append(namen, "	Anni	")
	namen = append(namen, "	Fabienne	")
	namen = append(namen, "	Zoey	")
	namen = append(namen, "	Sina	")
	namen = append(namen, "	Miriam	")
	namen = append(namen, "	Leila	")
	namen = append(namen, "	Linda	")
	namen = append(namen, "	Aylin	")
	namen = append(namen, "	Samira	")
	namen = append(namen, "	Elina	")
	namen = append(namen, "	Jolina	")
	namen = append(namen, "	Celina	")
	namen = append(namen, "	Elisabeth	")
	namen = append(namen, "	Valentina	")
	namen = append(namen, "	Julie	")
	namen = append(namen, "	Kira	")
	namen = append(namen, "	Alissa	")
	namen = append(namen, "	Olivia	")
	namen = append(namen, "	Jette	")
	namen = append(namen, "	Kim	")
	namen = append(namen, "	Elif	")
	namen = append(namen, "	Aaliyah	")
	namen = append(namen, "	Ela	")
	namen = append(namen, "	Lotte	")
	namen = append(namen, "	Anastasia	")
	namen = append(namen, "	Luana	")
	namen = append(namen, "	Hailey	")
	namen = append(namen, "	Lucia	")
	namen = append(namen, "	Lenja	")
	namen = append(namen, "	Rosalie	")
	namen = append(namen, "	Vivien	")
	namen = append(namen, "	Mona	")
	namen = append(namen, "	Lana	")
	namen = append(namen, "	Carolin	")
	namen = append(namen, "	Juna	")
	namen = append(namen, "	Elli	")
	namen = append(namen, "	Lynn	")
	namen = append(namen, "	Diana	")
	namen = append(namen, "	Thea	")
	namen = append(namen, "	Alexandra	")
	namen = append(namen, "	Angelina	")
	namen = append(namen, "	Carolina	")
	namen = append(namen, "	Marla	")
	namen = append(namen, "	Michelle	")
	namen = append(namen, "	Tessa	")
	namen = append(namen, "	Tabea	")
	namen = append(namen, "	Celine	")
	namen = append(namen, "	Leticia	")
	namen = append(namen, "	Svea	")
	namen = append(namen, "	Alisa	")
	namen = append(namen, "	Marleen	")
	namen = append(namen, "	Milla	")
	namen = append(namen, "	Amalia	")
	namen = append(namen, "	Joleen	")
	namen = append(namen, "	Mariella	")
	namen = append(namen, "	Laila	")
	namen = append(namen, "	Liana	")
	namen = append(namen, "	Rebecca	")
	namen = append(namen, "	Alessia	")
	namen = append(namen, "	Kimberly	")
	namen = append(namen, "	Nala	")
	namen = append(namen, "	Nelly	")
	namen = append(namen, "	Alicia	")
	namen = append(namen, "	Annalena	")
	namen = append(namen, "	Emmi	")
	namen = append(namen, "	Aurelia	")
	namen = append(namen, "	Lene	")
	namen = append(namen, "	Christina	")
	namen = append(namen, "	Samantha	")
	namen = append(namen, "	Larissa	")
	namen = append(namen, "	Noemi	")
	namen = append(namen, "	Dana	")
	namen = append(namen, "	Ina	")
	namen = append(namen, "	Evelyn	")
	namen = append(namen, "	Maira	")
	namen = append(namen, "	Anne	")
	namen = append(namen, "	Natalie	")
	namen = append(namen, "	Alma	")
	namen = append(namen, "	Amelia	")
	namen = append(namen, "	Giulia	")
	namen = append(namen, "	Lorena	")
	namen = append(namen, "	Fenja	")
	namen = append(namen, "	Zeynep	")
	namen = append(namen, "	Leona	")
	namen = append(namen, "	Tilda	")
	namen = append(namen, "	Felicitas	")
	namen = append(namen, "	Liv	")
	namen = append(namen, "	Liliana	")
	namen = append(namen, "	Nisa	")
	namen = append(namen, "	Veronika	")
	namen = append(namen, "	Jara	")
	namen = append(namen, "	Xenia	")
	namen = append(namen, "	Amira	")
	namen = append(namen, "	Linnea	")
	namen = append(namen, "	Medina	")
	namen = append(namen, "	Tuana	")
	namen = append(namen, "	Malia	")
	namen = append(namen, "	Henriette	")
	namen = append(namen, "	Jonna	")
	namen = append(namen, "	Jessika	")
	namen = append(namen, "	Cataleya	")
	namen = append(namen, "	Naila	")
	namen = append(namen, "	Valerie	")
	namen = append(namen, "	Alexa	")
	namen = append(namen, "	Carina	")
	namen = append(namen, "	Dilara	")
	namen = append(namen, "	Estelle	")
	namen = append(namen, "	Daria	")
	namen = append(namen, "	Joline	")
	namen = append(namen, "	Elise	")
	namen = append(namen, "	Helen	")
	namen = append(namen, "	Josie	")
	namen = append(namen, "	Rosa	")
	namen = append(namen, "	Azra	")
	namen = append(namen, "	Tamina	")
	namen = append(namen, "	Ava	")
	namen = append(namen, "	Enna	")
	namen = append(namen, "	Bella	")
	namen = append(namen, "	Leana	")
	namen = append(namen, "	Melanie	")
	namen = append(namen, "	Alena	")
	namen = append(namen, "	Cheyenne	")
	namen = append(namen, "	Enie	")
	namen = append(namen, "	Melia	")
	namen = append(namen, "	Meryem	")
	namen = append(namen, "	Esma	")
	namen = append(namen, "	Leandra	")
	namen = append(namen, "	Livia	")
	namen = append(namen, "	Selma	")
	namen = append(namen, "	Malin	")
	namen = append(namen, "	Nela	")
	namen = append(namen, "	Ylvi	")
	namen = append(namen, "	Ashley	")
	namen = append(namen, "	Madita	")
	namen = append(namen, "	Marina	")
	namen = append(namen, "	Marlena	")
	namen = append(namen, "	Janne	")
	namen = append(namen, "	Jill	")
	namen = append(namen, "	Maike	")
	namen = append(namen, "	Rieke	")
	namen = append(namen, "	Amina	")
	namen = append(namen, "	Ayla	")
	namen = append(namen, "	Melinda	")
	namen = append(namen, "	Alea	")
	namen = append(namen, "	Amilia	")
	namen = append(namen, "	Aurora	")
	namen = append(namen, "	Mailin	")
	namen = append(namen, "	Elin	")
	namen = append(namen, "	Enya	")
	namen = append(namen, "	Florentine	")
	namen = append(namen, "	Selin	")
	namen = append(namen, "	Valeria	")
	namen = append(namen, "	Annelie	")
	namen = append(namen, "	Heidi	")
	namen = append(namen, "	Malina	")
	namen = append(namen, "	Nicole	")
	namen = append(namen, "	Nika	")
	namen = append(namen, "	Flora	")
	namen = append(namen, "	Holly	")
	namen = append(namen, "	Liya	")
	namen = append(namen, "	Josefin	")
	namen = append(namen, "	Lenia	")
	namen = append(namen, "	Milana	")
	namen = append(namen, "	Tamara	")
	namen = append(namen, "	Asya	")
	namen = append(namen, "	Freya	")
	namen = append(namen, "	Lilian	")
	namen = append(namen, "	Talia	")
	namen = append(namen, "	Alice	")
	namen = append(namen, "	Mary	")
	namen = append(namen, "	Eliana	")
	namen = append(namen, "	Felina	")
	namen = append(namen, "	Hermine	")
	namen = append(namen, "	Mathea	")
	namen = append(namen, "	Sonja	")
	namen = append(namen, "	Alisha	")
	namen = append(namen, "	Soraya	")
	namen = append(namen, "	Elaine	")
	namen = append(namen, "	Madeleine	")
	namen = append(namen, "	Jolie	")
	namen = append(namen, "	Ceylin	")
	namen = append(namen, "	Eda	")
	namen = append(namen, "	Svenja	")
	namen = append(namen, "	Jamie	")
	namen = append(namen, "	Kate	")
	namen = append(namen, "	Lilith	")
	namen = append(namen, "	Madlen	")
	namen = append(namen, "	Mariam	")
	namen = append(namen, "	Miley	")
	namen = append(namen, "	Saskia	")
	namen = append(namen, "	Tiana	")
	namen = append(namen, "	Abby	")
	namen = append(namen, "	Aleyna	")
	namen = append(namen, "	Ann	")
	namen = append(namen, "	Edda	")
	namen = append(namen, "	Jolien	")
	namen = append(namen, "	Adriana	")
	namen = append(namen, "	Cara	")
	namen = append(namen, "	Hedi	")
	namen = append(namen, "	Ilayda	")
	namen = append(namen, "	Jenna	")
	namen = append(namen, "	Miray	")
	namen = append(namen, "	Alia	")
	namen = append(namen, "	Elsa	")
	namen = append(namen, "	Esila	")
	namen = append(namen, "	Jennifer	")
	namen = append(namen, "	Alexia	")
	namen = append(namen, "	Ellen	")
	namen = append(namen, "	Felicia	")
	namen = append(namen, "	Janina	")
	namen = append(namen, "	Joana	")
	namen = append(namen, "	Kaja	")
	namen = append(namen, "	Liara	")
	namen = append(namen, "	Marit	")
	namen = append(namen, "	Juliana	")
	namen = append(namen, "	Juliane	")
	namen = append(namen, "	Lilia	")
	namen = append(namen, "	Smilla	")
	namen = append(namen, "	Talea	")
	namen = append(namen, "	Viola	")
	namen = append(namen, "	Anouk	")
	namen = append(namen, "	Charlotta	")
	namen = append(namen, "	Jasmina	")
	namen = append(namen, "	Levke	")
	namen = append(namen, "	Aimee	")
	namen = append(namen, "	Ecrin	")
	namen = append(namen, "	Malea	")
	namen = append(namen, "	Marieke	")
	namen = append(namen, "	Naemi	")
	namen = append(namen, "	Mathilde	")
	namen = append(namen, "	Adelina	")
	namen = append(namen, "	Melek	")
	namen = append(namen, "	Melisa	")
	namen = append(namen, "	Naomi	")
	namen = append(namen, "	Nike	")
	namen = append(namen, "	Philine	")
	namen = append(namen, "	Shania	")
	namen = append(namen, "	Verena	")
	namen = append(namen, "	Cora	")
	namen = append(namen, "	Felia	")
	namen = append(namen, "	Malou	")
	namen = append(namen, "	Patricia	")
	namen = append(namen, "	Bianca	")
	namen = append(namen, "	Claire	")
	namen = append(namen, "	Delia	")
	namen = append(namen, "	Friederike	")
	namen = append(namen, "	Giuliana	")
	namen = append(namen, "	Yagmur	")
	namen = append(namen, "	Cassandra	")
	namen = append(namen, "	Joy	")
	namen = append(namen, "	Loreen	")
	namen = append(namen, "	Sena	")
	namen = append(namen, "	Tara	")
	namen = append(namen, "	Ceyda	")
	namen = append(namen, "	Eslem	")
	namen = append(namen, "	Helin	")
	namen = append(namen, "	Jona	")
	namen = append(namen, "	Lola	")
	namen = append(namen, "	Malena	")
	namen = append(namen, "	Melody	")
	namen = append(namen, "	Romina	")
	namen = append(namen, "	Anja	")
	namen = append(namen, "	Fatima	")
	namen = append(namen, "	Zara	")
	namen = append(namen, "	Zehra	")
	namen = append(namen, "	Annemarie	")
	namen = append(namen, "	Cecilia	")
	namen = append(namen, "	Dalia	")
	namen = append(namen, "	Elea	")
	namen = append(namen, "	Ellie	")
	namen = append(namen, "	Katja	")
	namen = append(namen, "	Melis	")
	namen = append(namen, "	Stefanie	")
	namen = append(namen, "	Tina	")
	namen = append(namen, "	Feyza	")
	namen = append(namen, "	Fine	")
	namen = append(namen, "	Josephina	")
	namen = append(namen, "	Vivian	")
	namen = append(namen, "	Adele	")
	namen = append(namen, "	Alva	")
	namen = append(namen, "	Eleni	")
	namen = append(namen, "	Eliza	")
	namen = append(namen, "	Enni	")
	namen = append(namen, "	Franka	")
	namen = append(namen, "	Janna	")
	namen = append(namen, "	Maileen	")
	namen = append(namen, "	Maxi	")
	namen = append(namen, "	Sidney	")
	namen = append(namen, "	Ada	")
	namen = append(namen, "	Amara	")
	namen = append(namen, "	Inga	")
	namen = append(namen, "	Leia	")
	namen = append(namen, "	Liz	")
	namen = append(namen, "	Lou	")
	namen = append(namen, "	Lydia	")
	namen = append(namen, "	Marisa	")
	namen = append(namen, "	Sila	")
	namen = append(namen, "	Stina	")
	namen = append(namen, "	Tamia	")
	namen = append(namen, "	Alara	")
	namen = append(namen, "	Anisa	")
	namen = append(namen, "	Cleo	")
	namen = append(namen, "	Megan	")
	namen = append(namen, "	Nea	")
	namen = append(namen, "	Penelope	")
	namen = append(namen, "	Zümra	")
	namen = append(namen, "	Beyza	")
	namen = append(namen, "	Charleen	")
	namen = append(namen, "	Femke	")
	namen = append(namen, "	Henrieke	")
	namen = append(namen, "	Jamila	")
	namen = append(namen, "	Jenny	")
	namen = append(namen, "	Mirja	")
	namen = append(namen, "	Nila	")
	namen = append(namen, "	Salome	")
	namen = append(namen, "	Sandra	")
	namen = append(namen, "	Alessa	")
	namen = append(namen, "	Christin	")
	namen = append(namen, "	Evelina	")
	namen = append(namen, "	Joyce	")
	namen = append(namen, "	Kiana	")
	namen = append(namen, "	Line	")
	namen = append(namen, "	Natalia	")
	namen = append(namen, "	Ria	")
	namen = append(namen, "	Tanja	")
	namen = append(namen, "	Betty	")
	namen = append(namen, "	Davina	")
	namen = append(namen, "	Defne	")
	namen = append(namen, "	Denise	")
	namen = append(namen, "	Dila	")
	namen = append(namen, "	Eleonora	")
	namen = append(namen, "	Gloria	")
	namen = append(namen, "	Judith	")
	namen = append(namen, "	Julika	")
	namen = append(namen, "	Käthe	")
	namen = append(namen, "	Katrin	")
	namen = append(namen, "	Laureen	")
	namen = append(namen, "	Leonora	")
	namen = append(namen, "	Lisbeth	")
	namen = append(namen, "	Luzi	")
	namen = append(namen, "	Maxima	")
	namen = append(namen, "	Neyla	")
	namen = append(namen, "	Nisanur	")
	namen = append(namen, "	Phoebe	")
	namen = append(namen, "	Ruby	")
	namen = append(namen, "	Sabrina	")
	namen = append(namen, "	Vera	")
	namen = append(namen, "	Ziva	")
	namen = append(namen, "	Abigail	")
	namen = append(namen, "	Alya	")
	namen = append(namen, "	Andrea	")
	namen = append(namen, "	Ariana	")
	namen = append(namen, "	Belinay	")
	namen = append(namen, "	Fanny	")
	namen = append(namen, "	Ivy	")
	namen = append(namen, "	Joanna	")
	namen = append(namen, "	Jolin	")
	namen = append(namen, "	Lavin	")
	namen = append(namen, "	Maren	")
	namen = append(namen, "	Melike	")
	namen = append(namen, "	Nia	")
	namen = append(namen, "	Nova	")
	namen = append(namen, "	Saphira	")
	namen = append(namen, "	Tia	")
	namen = append(namen, "	Amanda	")
	namen = append(namen, "	Ariane	")
	namen = append(namen, "	Arina	")
	namen = append(namen, "	Dorothea	")
	namen = append(namen, "	Emina	")
	namen = append(namen, "	Feline	")
	namen = append(namen, "	Julina	")
	namen = append(namen, "	Lieselotte	")
	namen = append(namen, "	Luca	")
	namen = append(namen, "	Mette	")
	namen = append(namen, "	Narin	")
	namen = append(namen, "	Nilay	")
	namen = append(namen, "	Philippa	")
	namen = append(namen, "	Polly	")
	namen = append(namen, "	Rafaela	")
	namen = append(namen, "	Ruth	")
	namen = append(namen, "	Sharon	")
	namen = append(namen, "	Summer	")
	namen = append(namen, "	Clarissa	")
	namen = append(namen, "	Elanur	")
	namen = append(namen, "	Esther	")
	namen = append(namen, "	Isa	")
	namen = append(namen, "	Liyana	")
	namen = append(namen, "	Nadine	")
	namen = append(namen, "	Sarina	")
	namen = append(namen, "	Serafina	")
	namen = append(namen, "	Violetta	")
	namen = append(namen, "	Yasmina	")
	namen = append(namen, "	Ylva	")
	namen = append(namen, "	Acelya	")
	namen = append(namen, "	Anita	")
	namen = append(namen, "	Annabella	")
	namen = append(namen, "	Ceren	")
	namen = append(namen, "	Damla	")
	namen = append(namen, "	Fatma	")
	namen = append(namen, "	Fina	")
	namen = append(namen, "	Frederike	")
	namen = append(namen, "	Grace	")
	namen = append(namen, "	Lale	")
	namen = append(namen, "	Leevke	")
	namen = append(namen, "	Mareike	")
	namen = append(namen, "	Mieke	")
	namen = append(namen, "	Rahel	")
	namen = append(namen, "	Stine	")
	namen = append(namen, "	Timea	")
	namen = append(namen, "	Wiebke	")
	namen = append(namen, "	Alicja	")
	namen = append(namen, "	Anneke	")
	namen = append(namen, "	Mia	")
	namen = append(namen, "	Emma	")

}