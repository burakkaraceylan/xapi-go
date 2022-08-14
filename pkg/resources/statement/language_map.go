package statement

// A language map is a dictionary where the key is a RFC 5646 Language Tag, and the value is a string in the language specified in the tag.
// This map SHOULD be populated as fully as possible based on the knowledge of the string in question in different languages.
// https://github.com/adlnet/xAPI-Spec/blob/master/xAPI-Data.md#42-language-maps
type LanguageMap map[string]string
