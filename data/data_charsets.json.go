// This file is automatically generated by generate-charset-data.
// Do not hand-edit.

package data

import (
	"io"
	"io/ioutil"
	"strings"

	"github.com/justwatchcom/go-charset/charset"
)

func init() {
	charset.RegisterDataFile("charsets.json", func() (io.ReadCloser, error) {
		r := strings.NewReader("{\n\"8bit\": {\n\t\"Desc\": \"raw 8-bit data\",\n\t\"Class\": \"8bit\",\n\t\"Comment\": \"special class for raw 8bit data that has been converted to utf-8\"\n},\n\"big5\": {\n\t\"Desc\": \"Big 5 (HKU)\",\n\t\"Class\": \"big5\",\n\t\"Comment\": \"Traditional Chinese\"\n},\n\"euc-jp\": {\n\t\"Aliases\":[\"x-euc-jp\"],\n\t\"Desc\": \"Japanese Extended UNIX Code\",\n\t\"Class\": \"euc-jp\"\n},\n\"gb2312\": {\n\t\"Aliases\":[\"iso-ir-58\", \"chinese\", \"gb_2312-80\"],\n\t\"Desc\": \"Chinese mixed one byte\",\n\t\"Class\": \"gb2312\"\n},\n\"ibm437\": {\n\t\"Aliases\":[\"437\", \"cp437\"],\n\t\"Desc\": \"IBM PC: CP 437\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"ibm437.cp\",\n\t\"Comment\": \"originally from jhelling@cs.ruu.nl (Jeroen Hellingman)\"\n},\n\"ibm850\": {\n\t\"Aliases\":[\"850\", \"cp850\"],\n\t\"Desc\": \"IBM PS/2: CP 850\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"ibm850.cp\",\n\t\"Comment\": \"originally from jhelling@cs.ruu.nl (Jeroen Hellingman)\"\n},\n\"ibm866\": {\n\t\"Aliases\":[\"cp866\", \"866\"],\n\t\"Desc\": \"Russian MS-DOS CP 866\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"ibm866.cp\"\n},\n\"iso-8859-1\": {\n\t\"Aliases\":[\"iso-ir-100\", \"ibm819\", \"l1\", \"iso8859-1\", \"iso-latin-1\", \"iso_8859-1:1987\", \"cp819\", \"iso_8859-1\", \"iso8859_1\", \"latin1\"],\n\t\"Desc\": \"Latin-1\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"iso-8859-1.cp\"\n},\n\"iso-8859-10\": {\n\t\"Aliases\":[\"iso_8859-10:1992\", \"l6\", \"iso-ir-157\", \"latin6\"],\n\t\"Desc\": \"Latin-6\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"iso-8859-10.cp\",\n\t\"Comment\": \"originally from dkuug.dk:i18n/charmaps/ISO_8859-10:1993\"\n},\n\"iso-8859-15\": {\n\t\"Aliases\":[\"l9-iso-8859-15\", \"latin9\"],\n\t\"Desc\": \"Latin-9\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"iso-8859-15.cp\"\n},\n\"iso-8859-2\": {\n\t\"Aliases\":[\"iso-ir-101\", \"iso_8859-2:1987\", \"l2\", \"iso_8859-2\", \"latin2\"],\n\t\"Desc\": \"Latin-2\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"iso-8859-2.cp\"\n},\n\"iso-8859-3\": {\n\t\"Aliases\":[\"iso-ir-109\", \"l3\", \"iso_8859-3:1988\", \"iso_8859-3\", \"latin3\"],\n\t\"Desc\": \"Latin-3\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"iso-8859-3.cp\"\n},\n\"iso-8859-4\": {\n\t\"Aliases\":[\"iso-ir-110\", \"iso_8859-4:1988\", \"l4\", \"iso_8859-4\", \"latin4\"],\n\t\"Desc\": \"Latin-4\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"iso-8859-4.cp\"\n},\n\"iso-8859-5\": {\n\t\"Aliases\":[\"cyrillic\", \"iso_8859-5\", \"iso-ir-144\", \"iso_8859-5:1988\"],\n\t\"Desc\": \"Part 5 (Cyrillic)\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"iso-8859-5.cp\"\n},\n\"iso-8859-6\": {\n\t\"Aliases\":[\"ecma-114\", \"iso_8859-6:1987\", \"arabic\", \"iso_8859-6\", \"asmo-708\", \"iso-ir-127\"],\n\t\"Desc\": \"Part 6 (Arabic)\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"iso-8859-6.cp\"\n},\n\"iso-8859-7\": {\n\t\"Aliases\":[\"greek8\", \"elot_928\", \"ecma-118\", \"greek\", \"iso_8859-7\", \"iso_8859-7:1987\", \"iso-ir-126\"],\n\t\"Desc\": \"Part 7 (Greek)\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"iso-8859-7.cp\"\n},\n\"iso-8859-8\": {\n\t\"Aliases\":[\"iso_8859-8:1988\", \"hebrew\", \"iso_8859-8\", \"iso-ir-138\"],\n\t\"Desc\": \"Part 8 (Hebrew)\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"iso-8859-8.cp\"\n},\n\"iso-8859-9\": {\n\t\"Aliases\":[\"l5\", \"iso_8859-9:1989\", \"iso_8859-9\", \"iso-ir-148\", \"latin5\"],\n\t\"Desc\": \"Latin-5\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"iso-8859-9.cp\"\n},\n\"koi8-r\": {\n\t\"Desc\": \"KOI8-R (RFC1489)\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"koi8-r.cp\"\n},\n\"shift_jis\": {\n\t\"Aliases\":[\"sjis\", \"ms_kanji\", \"x-sjis\"],\n\t\"Desc\": \"Shift-JIS Japanese\",\n\t\"Class\": \"cp932\",\n\t\"Arg\": \"shiftjis\"\n},\n\"utf-16\": {\n\t\"Aliases\":[\"utf16\"],\n\t\"Desc\": \"Unicode UTF-16\",\n\t\"Class\": \"utf16\"\n},\n\"utf-16be\": {\n\t\"Aliases\":[\"utf16be\"],\n\t\"Desc\": \"Unicode UTF-16 big endian\",\n\t\"Class\": \"utf16\",\n\t\"Arg\": \"be\"\n},\n\"utf-16le\": {\n\t\"Aliases\":[\"utf16le\"],\n\t\"Desc\": \"Unicode UTF-16 little endian\",\n\t\"Class\": \"utf16\",\n\t\"Arg\": \"le\"\n},\n\"utf-8\": {\n\t\"Aliases\":[\"utf8\", \"ascii\", \"us-ascii\"],\n\t\"Desc\": \"Unicode UTF-8\",\n\t\"Class\": \"utf8\"\n},\n\"windows-1250\": {\n\t\"Desc\": \"MS Windows CP 1250 (Central Europe)\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"windows-1250.cp\"\n},\n\"windows-1251\": {\n\t\"Desc\": \"MS Windows CP 1251 (Cyrillic)\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"windows-1251.cp\"\n},\n\"windows-1252\": {\n\t\"Desc\": \"MS Windows CP 1252 (Latin 1)\",\n\t\"Class\": \"cp\",\n\t\"Arg\": \"windows-1252.cp\"\n},\n\"windows-31j\": {\n\t\"Aliases\":[\"cp932\"],\n\t\"Desc\": \"MS-Windows Japanese (cp932)\",\n\t\"Class\": \"cp932\",\n\t\"Arg\": \"cp932\"\n}\n}\n")
		return ioutil.NopCloser(r), nil
	})
}
