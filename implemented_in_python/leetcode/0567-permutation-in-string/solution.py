# https://leetcode.com/problems/permutation-in-string/


class SimplerSolution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        s1Count = {}

        for l in s1:
            s1Count[l] = s1Count.get(l, 0) + 1

        s2Count = {}
        left, right = 0, 0
        while right < len(s2):
            rightLetter = s2[right]
            s2Count[rightLetter] = s2Count.get(rightLetter, 0) + 1

            if right - left + 1 < len(s1):
                right += 1
                continue

            if self.areEqual(s1Count, s2Count):
                return True

            leftLetter = s2[left]
            s2Count[leftLetter] -= 1
            left += 1
            right += 1

        return False

    def areEqual(self, s1Count, s2Count):
        for letter, count in s1Count.items():
            c = s2Count.get(letter, 0)
            if c != count:
                return False
        return True


class Solution:
    # O(s2) + O(s1) => O(s2)
    def checkInclusion(self, s1: str, s2: str) -> bool:
        letterCount = {}
        for l in s1:
            letterCount[l] = 1 + letterCount.get(l, 0)

        left, right = 0, 0
        while right < len(s2):
            rightLetter = s2[right]
            count = letterCount.get(rightLetter, 0)

            # not found
            if count == 0:
                leftLetter = s2[left]
                left += 1

                if leftLetter in letterCount:
                    letterCount[leftLetter] += 1
                    continue

                right += 1
                continue

            # check length
            if right - left + 1 == len(s1):
                return True

            letterCount[rightLetter] -= 1
            right += 1

        return False


# Neetcode solution
class NeetcodeSolution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False

        s1Count, s2Count = [0] * 26, [0] * 26
        for i in range(len(s1)):
            s1Count[ord(s1[i]) - ord("a")] += 1
            s2Count[ord(s2[i]) - ord("a")] += 1

        matches = 0
        for i in range(26):
            matches += 1 if s1Count[i] == s2Count[i] else 0

        l = 0
        for r in range(len(s1), len(s2)):
            if matches == 26:
                return True

            index = ord(s2[r]) - ord("a")
            s2Count[index] += 1
            if s1Count[index] == s2Count[index]:
                matches += 1
            elif s1Count[index] + 1 == s2Count[index]:
                matches -= 1

            index = ord(s2[l]) - ord("a")
            s2Count[index] -= 1
            if s1Count[index] == s2Count[index]:
                matches += 1
            elif s1Count[index] - 1 == s2Count[index]:
                matches -= 1
            l += 1
        return matches == 26


sol = Solution()

s1 = "cgzvjlgigiosbizrtoyxexeijhqzzgbwxnmyhmanpaagjblooflawusyjiayocqhpxkybgkkybkostijmntcaxpjrfrvirvwdvlcrnrdsceldnhihrftexsdykkzcbelecdjhnwfuvnwwyxvowugdhpxfbigotfajpitxuvvcmbiipjzctrrlyxswgcvkaklfkocmathfdsgddmtstzqedazztmutaqgyiywsufcgejubzsuwcawpxwwdnzfvkpbdjuvhkrifzdtzpssbvqetzejjpdfzxczfnfoucfkvyspmpscpljmchrsfdtfcyvhfyqiaxfqdmewxtwavbxreeogxeelvrashuelrkrmxmdxriayxfmhfmxelwbydycfornyqpuhjrbgbukotgtgbtlnqljjxktnysnayiutkxuzdimbaocjymgledptwavfhlnauiafqiaoyngspnmttcxkgzufpvfvfulclgqekmehwqmfqtcrtvkfnjifrhmrsbcixlopgxkrgwnqjdhrftrwqmnfgjhwmbbmbnuawyjliarjgafybqvcigmpiyqguiwzvthtwqjvwmgfxvfwcyxizavpwwwuoqzhjyjxnjxyzpdmajxtckzdkklipdldzobgjnhyrcexggtjbilqrnfgvmtwrvhxsuysxyuumnxaucvkuafpbegihgcsuhfqqljrxrquaqpjmsegzisqogtvxcqfhnrxedovyjgyknqlfmdlpfbcyfrdokvlzlgbsnajwgnybmutvljqwovppctbadmkmujpwnggqbrcqmjbdaodixjsozzrmgoycpfkhwomujbuxqdyjmrqbchhfpwgbigboxseknxdptgoejsgewykoszssnnwaocaodqsghprzqbtnfosgotvdwxzwhjdlqhlwgidlrdudgwkgrtzdxjlhbkdjkjxjvmkvdimdeejvsmqjkagffrfsdspstowzcablbtcbpdovfinpyatkgcrzotmtztiyqnxycsawmdpfdeubekhzktgnqabfhqkfrszwpaylarurklzrfjeruqpjhhjibojcovcbccbscgwwjkixhufvbmrbbcbjiaxhhryqgbbpbzwbbbyxukrervckfislgwunqbyzandfqlsturgrssnzneqtgbsrfugstirqholqmftjacnqapjdxzrcpqfnkiuaqeyawhobemdnkshevtkywueowgekedmjdledcddtoomnkkzgvyqbonynqlakhcxynsfoxjlkevqscltkyroiiymzcxvmvhtojspzwezdxabuxuhhrgeynczmnujzglzhgywtpbyygwvazmqbcnlbjwcyzkurnsrvuxltsymrvoonasiquqvtgcffaphsfklonlsdvrmmobdafchuyntcciugpfgkuszgwrcfkbpfztiwpcczlnjoabxvtkldmtyjlbhorvhiuvzqidypyzwxqjefdhztbgcunthxzimaszfvzzfbnydxsmnplqgtrkcbadmvbqlonrxdgxpunwptatdvgxnybmviiwdrkcplezholphdenmvywilaqfujnmbeydvovytlmzpjalkekwbljwvrmemfcthyickamqdszdglrsapjjhehbiqoxhgynfyypljxhpfjxlckfxvhetcylewtnbhvgvhbpsmftgzsrzawwcunjdarwdgsqrihyshnraiizbsavcdqmjvezehjjgaxnmffdtapkzbhimygljaimqncsrynjzvkdcenndicaxnorrcbhvkgrouchpzqxtchmqabktglnjbpxhuroxjxskfppkyzbsvkenecglzbgvkxdpkjoykcsobkjnzwhbyycvkxwdwqbfzjeoakbadlaekmqgdxbojwpvkftckuparzpeyytdowcxrupeoidlirdqseqhkpfxcvfuuyqybgosfppscdtwetliojdztckjarmlkmtpuhqfhshdwummaphq"
s2 = "vnndchsddonunovgxxlrrdtnfjidmocdrdcnousrajalpsgvxolzhujffwszxhqhojtmdluwdrzwqwqvbebztgghadtdroxujmzwmftiwamymwmpysqwffjuxtdebqoglbqobxfixdlircftggnezxyxrjzytffudihaespkhdpeqnalqohjqtypsobhrxtovxskqvtdqulgaeozlxazpubdveorkytajakzagxmhrlcvfwyuteboazjnecgrggpmmxplndhbvkhmkaozlqpgxenkchjjtaegwfcbquxlsldlbxzyfwahdpybheugrysnbjizbpgpwzwfdnriscqlrndkvibkqczjbiwkahypuujuexwtbtiqpnixwcslojlowxxjhpxvmxssciyhxpbmfpoeiizuhlkswctwkzguwyvyqvtemnerodalfuocyaduvyjwcgyxtcdhggskcdazjsqnowkkklmnocgrjpgrwdwjbhvvwtgfuialbafnycigyhnphhvnibinvadcoprwdnreuwukosnlatsymuqxtgdvsrvmercfpfaaeszkhknurpqxhyyxclzvzjqkztoqrpszcmhbcdvxwbdokmflzcsepejdemzvtzhtkmhhjhducxxijvgjidjaemzwkeyggircxivxyxcogasotyfnuamfavpkcjbnvxddhyiiybwkqobmzuzmegdiqvzjptiwhujpndvxzlaclqsemhtqwufeonsdjgnbkkjjsyxtpdytfsgpmknjibjyxctwmprbcyjhlfiyaofvjryiicpoaxeonayvztnkvzpoprnbdtwllozlqyedvnezxxxuspenwxjekkgaeajhxriahqcutakhifqkzgcjldlkbxapqrwvwshpyvnbukblgcdyyhikloslaxrsfrkwexzbrndopgnwejltpdiiwykwfymmwtsbjcvtmnakjyeqcbcbdtcuxuueaxzurixojxqdfcsdmfejajqlissfkkowchuyabccuazveqmegrtdnsoqsqxgmilxnyewpkiheajdvcnyjuuuvhonjumnshoinriazaxrqizrkuhyadjmkwdizkuamgseqibgnhgovccxmgtcywroxctticxbmvvqenvmuzdobfcxqffzgcojdhvvvkgajmuelevryxxjcvfbgpztrpzjpmttwjalllbltzmfqiwoodlvnnoviwiotmpjkjwfdnjfcsylhtioszxytvkkhrunwqxuivzukynbrdjryxkivngsecgskxhtpffrvfpwrpcythjbjdjocxcaemznunyhhrvcchmubozsqnfbgwrtzhhaprjkiwwrxzggkamvzdeidbckwyctuspdmfjhhrujobvvzuhhimutzlqfleimwechdemrztldrqwtevkbmttsdxcpmkzogegytvtrwevzbvszsxpohqqndstnstnaltilpeqdixpokcgzvjlgigiosbizrtoyxexeijkqzzgbwxnmyhmanpmagjblooflawusyjiayocqhpxkpbgkkybkostijmntcaxpjrfrvirvwdvlcrnrdsceldnhihrfyexsdykkzcbelecdjhnqcuvnwwyxvowugdhpxfbigotfajpitxuvvcmbiipjzctrrlyxswgcvkahlfkocmathfdsgddmtstzqedazztmutaqgyitwsufcgejufzsuwcawpxwwdnzfvkpbdjuvhkrifzdtzpssbvqetzejjpdfzxczfnfoucfkvyspmpscpljmchrsfdqfcyvhfyqiaxfqdmewxtwavbxreeogxeelvrashuelrkrmxmdxriayxfmhfmxelwbydycfornyqpuhjrbgbukotgtgbtlnqljjxktnysnayiutkxuzdimbsocjyagledptwavfhlnauiafqiaoyngspnmttcxkgzubpvfvfulclgqekmehwqmfqtcrtxkfnjifrhmrsbcixlougxkrgwnqjdhrftrwqmnfgjhwmbbmbnuawyjliarjgafylqvcigmpiyqguiwzvthtwqjvwmgfxvfwcyvizavpwwwuoqzhjyjxnjxyzpdmajxtckzdkklipdldzobgjnhyrcexggtjbilqrnfgvmtwrvhxsuysxyuumnxaucvkuafpbegihgcsuhfqqljrxrquaqpjmsegzisqogtvxcqfhnrxedovyjgykntlfmdlpfbcyfrdokvlzlgbsnajwgnybmutvljqwovppctbadmkmujpwnggqbrcqmjbdaodixjsozzrmgoyfpfkhwomajbuxqdyjmrqbchhfpwgbigboxseknxdptgoejsgewykoszssnywaocaodqsghprzqbtnfosgotvdwxuwhjdlqhlwgidbrdudgwkgrtzdxjlhbkdjkjxjvmkvdimdeejvsmqjkagffrfsdspstowzcablbtcbpdovfinpyutkgcrzotmtztiyqnxycsawmdpfdezbekhzktgnqabfhqkfrszwpaylarurklzrfjeruqpjhhjibojcovcbccbscgwwjkixhpfvbmrbbcbjiaxhhryqgbbpbzwbbbyxukrervckfislgwunqbyzandfqlsturgrssnzneqtgbsrfugstirqholqmftjacnqapjdxzrcpqfnkiuaqeyawhobemdnkshevtkywueowgekedmjdledcddtoomnkkzgvnqbonynqlakkcxynsfoxjlkevqsyltkyroiiymzcxvmvhtojspzwezdxabuxuhhrgecnczmnujzglzhgywtpbyygwvazmqbcnlbjwcyzkurnsrvuxltsymrvoonasiquqvtgcffaphsfklonlsdvrmmobdafchuyntcciugpfgkuszgwrcfkbpfztiwpcczlnjoabxvtkldmtyjlbhorvhiuvzqidypyzwxqjefdhztbgcunthxzimaszfvzzfbnydxsmnplqgtrkcbadmvbqlonrxdgxpunwptatdvgxnybmviiwdrkcplezholphdenmvywilaqfujnmbeydvovytlmzpjalkekwbljwvrmemfcthyickamqdszdglrsapjjhehbiqoxhgynfyypljxhpfjxlckfxvhetcylewtnbhvgvhbpsmftgzsrzawwcunjdarwdgaqrihyshnraiizbsavcdqmjvezehjjgaxnmrfdtapkzbhimygljaimqncsrynjzvkdcenndicaxnorrcbhvkgrouchpzqxtchmqabktglnjbpxhuroxjxskfppkyzbsvkenecglzbgvkxdykjoykcsobkjnzwhbyycvkxwdwqbfzjeoakbadlaehmwgdxbojwpvkftckuparzpeyytdowcxrupeoidlirdqseqhkpfxcvfuuyqybgosfppscdtwetliojdztckjafmlkmtpuhqfhshdwummaphqccfbvtdzgqhgxefzzpggjsilrfchvpzmkmmdxncnqiauqkpxldmynhhqceijcmucekiqtojnotkaqrjpdhxdpxhtabmxnszocainqyzuciyktazyksdvrcerptpjrbkszgehypijcqmvpezufyhscxgowkuhexfikfqdkxemkkowkfofxskwyumvmfvpparszcldalecfmkltqmxubmbmbnanciofqaxxz"
print(sol.checkInclusion(s1, s2))
