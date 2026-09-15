package lc2472

import (
	"reflect"
	"testing"
)

func TestMaxPalindromes(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{

		{
			name: "example 1",
			s:    "abaccdbbd",
			k:    3,
			want: 2,
		},

		{
			name: "example 2",
			s:    "adbcda",
			k:    2,
			want: 0,
		},

		{
			name: "example 3",
			s:    "abaaaaab",
			k:    3,
			want: 2,
		},

		{
			name: "example 4",
			s:    "fojnrcnerybjjghhaykxykixkojokxikyxkyahhgjjbyrencrnjofzyzeftizdmiynofcwlwcfonyimdzitfezyzqywfmxqcqhwbsmtbfberkuiukrebfbtmsbwhqcqxmfwywhoppuqnxqesiqedlpujohqpnvzvljyjlvzvnpqhojupldeqiseqxnquppohwrniimkadiwfkesjhhjsekfwidakmiingbsasadxeqxacdhxqkupyprzrcjaqtqajcrzrpypukqxhdcaxqexdasawuumyhklbhodwcmydtkfezcuesnjjnseuczefktdymcwdohblkhymuuwsdquovgnrckcpubplmueyaernzdmcmdznreayeumlpbupckcrngvouqdpsktcrqynclzrgjtplokkfihhymmgsgmmyhhifkkolptjgrzlcnyqrctksvnchfxupvsyljxzadjhohjdazxjlysvpuxfhcnwatcqotbwunztyzryjmmjyrzytznuwbtoqctawdykmoeebpeqmgbdexfwxcelvwxyrryxwvlecxwfxedbgmqepbeeblxevnmdcnrwfsaqqasfwrncdmnvexlbwaurakgpqtjpymaibpebeqlelqebepbiamypjtqpgkaruawceerqumnqbdwnbtlcootttibpdjpckqfcrmajjamrcfqkcpjdpbitttoocltbnwdbqnmuqreeyvfwueuewqyxyiaquqaiyxyqweueuwfvksskyuswbeqpihumbedwoxhufidbdifuhxowdebmuhipqebwsuykssktegrvozxyiofngmqbkiqfagxsfyllzgumemugzllyfsxgafqikbqmgnfoiyxzosthtycrjroeipdirvtrlzukccpqpcckuzlrtvridpieorjrcythtsdalhtjojbryxaoxppdxpwwbbklywnamrxtjpwtqfxyegmywwvwwymgeyxfqtwpjtxrmanwylkbbwwpxdppxoabymjvqlwqnmfzerregerrezfmnqwlqvjmybgogqdfigragpafjcyfzkqxzladjqufrrfuqjdalzxqkzfycjfapgargifdqgogarlptpazebeudrmymnvimxkssysytgujznjancjrjcnajnzjugtysysskxmivnmymrduebezaptphpmeblrzbxzzvxpwwxhiykxmbdcrwvcketekcvwrcdbmxkyihxwwpxvzzxbzrlbemhvcaswtkxbsdpalkxkddqokerskptsustbcbtsustpksrekoqddkxklapdsbxktwsacvhmjigkwrygxusxbbguwzngqoqbknehaopuvvupoahenkbqoqgnzwugbbxsuxgyrwkgikjureyadreoxreinabllbanierxoerdayerujkmjuxcoydsukgdctdcnrsqxcrzvmxmvzrcxqsrncdtcdgkusdyocxuuznoynsubjvoibuxfddwfogofwddfxubiovjbusnyonzuuubdw",
			k:    24,
			want: 29,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := maxPalindromes(tt.s, tt.k)

			if !reflect.DeepEqual(out, tt.want) {
				t.Fatalf("got %v want %v", out, tt.want)
			}
		})
	}
}
