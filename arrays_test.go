package main;

import "testing";

func TestIsUnique(t *testing.T) {
    tests := make(map[string]bool);

    tests["bye"] = true;
    tests[""] = true;
    tests[" "] = true;
    tests["abcdefghijklmnop"] = true;

    tests["aa"] = false;
    tests["hello"] = false; 
    tests["aaa"] = false;
    tests["aba"] = false;
    tests[" a "] = false;

    for input, want := range tests {
        res := IsUnique(input);
        if want != res {
            t.Fatalf("%s failed: wanted %t, received: %t", input, want, res);
        }
    }
}

func TestCheckPermutation(t *testing.T) {
    tests := make(map[[2]string]bool);

    tests[[2]string{"aba", "aba"}] = true;
    tests[[2]string{"aba", "baa"}] = true;
    tests[[2]string{"abcd", "dcba"}] = true;
    tests[[2]string{"pog", "gop"}] = true;
    tests[[2]string{"", ""}] = true;

    tests[[2]string{"aba", "aa"}] = false;
    tests[[2]string{"aba", "bab"}] = false;
    tests[[2]string{"", "a"}] = false;
    tests[[2]string{"abcdcd", "jfkdlf"}] = false;

    for input, want := range tests {
        res := CheckPermutation(input[0], input[1]);
        if want != res {
            t.Fatalf("%s - %s failed: wanted %t, received %t", input[0], input[1], want, res);
        }
    }
}

func TestUrlify(t *testing.T) {
    tests := make(map[string]string);

    tests["hello"] = "hello";
    tests["hel lo"] = "hel%20lo";
    tests["hello "] = "hello%20";
    tests[" hello"] = "%20hello";

    for key, want := range tests {
        res := Urlify(key);
        if want != res {
            t.Fatalf("%s failed: wanted %s, received %s", key, want, res);
        }
    }
}

func TestIsPermOfPalindrome(t *testing.T) {
    tests := make(map[string]bool);

    tests["racecar"] = true;
    tests["a"] = true;
    tests[""] = true;
    tests["aba"] = true;
    tests["aa"] = true;
    tests["aab"] = true;

    tests["ab"] = false;
    tests["abcdefg"] = false;

    for key, want := range tests {
        res := IsPermOfPalnidrome(key);
        if want != res {
            t.Fatalf("%s failed: wanted %t, received %t", key, want, res);
        }
    }
}
