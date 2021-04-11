/**
 * @param {string} S
 * @return {string}
 */
function isLetter(c) {
    return ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z');
}

var reverseOnlyLetters = function(S) {
    let answer = '';
    const dashs = {};
    let k = S.length-1;
    for (let i=0; i<S.length; i++) {
        const c = S[i];
        if (isLetter(c)) {
            while(!isLetter(S[k])) {
                k--;
            }
            answer += S[k];
            k--;
        } else {
            answer += c;
        }
    }
    return answer
};
