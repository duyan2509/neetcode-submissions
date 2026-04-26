func isAnagram(s string, t string) bool {
    if(len(s)!=len(t)){
        return false;
    }
    checkMap := make(map[byte]int)
    for i:=0;i < len(s); i++ {
        checkMap[s[i]]++;
    }
    for i:=0;i < len(t); i++{
        checkMap[t[i]]--;
        if checkMap[t[i]]<0 {
            return false;
        }
    }
    for _, value := range checkMap {
        if value!=0 {
            return false;
        }
    }
    return true;
}
