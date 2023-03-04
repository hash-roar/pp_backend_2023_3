package blockwordsservice

import "pp-bakcend/models"

func AddOrUpdate(mid string, word string, handle string) error {
	if models.CheckBlockWordExist(mid, word) {
		return models.UpdateBlockWord(mid, word, handle)
	}
	return models.AddBlockWords(mid, word, handle)
}

func Delete(mid string, word string) error {
	return models.DeleteBlockWord(mid, word)
}

func GetAllBlockWords() ([]models.BlockWordsInfo, error) {
	return models.GetAllBlockWordsInfo()
}

func SetWordVisibility(mid string, word string, visible bool) error {
	return models.SetWordVisibility(mid, word, visible)
}

type BlockWordsInfo struct {
	Mid    string
	Shield string
	Handle string
}

func GetUserBlockWords(mid string) ([]BlockWordsInfo, error) {
	var infos []BlockWordsInfo
	blockinfos, err := models.GetUserBlockWords(mid)
	if err != nil {
		return nil, err
	}
	for _, v := range blockinfos {
		infos = append(infos, BlockWordsInfo{
			Mid:    v.Mid,
			Shield: v.Word,
			Handle: v.Handle,
		})
	}
	return infos, nil
}
func GetBlockWordsWithoutRepetition(mid string) ([]BlockWordsInfo, error) {
	user_words, err := models.GetUserBlockWords(mid)
	if err != nil {
		return nil, err
	}
	all_words, err := models.GetAllBlockWordsInfo()
	if err != nil {
		return nil, err
	}
	var filter map[string]BlockWordsInfo
	for _, uword := range user_words {
		filter[uword.Word] = BlockWordsInfo{
			Mid:    uword.Mid,
			Shield: uword.Word,
			Handle: uword.Handle,
		}
	}
	for _, aword := range all_words {
		_, ok := filter[aword.Shield]
		if !ok && aword.Visible {
			filter[aword.Shield] = BlockWordsInfo{
				Mid:    aword.Mid,
				Shield: aword.Shield,
				Handle: aword.Handle,
			}
		}
	}

	infos := make([]BlockWordsInfo, len(filter))
	for _, v := range filter {
		infos = append(infos, v)
	}

	return infos, nil

}
