package app

type webComic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func (wc *webComic) write() error {
	db, err := dbConnect()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`
		INSERT INTO xkcd (month, num, link, year, news, safe_title, transcript,
		alt, img, title, day) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		ON CONFLICT ON CONSTRAINT xkcd_num_key DO NOTHING`,
		wc.Month, wc.Num, wc.Link, wc.Year, wc.News, wc.SafeTitle, wc.Transcript,
		wc.Alt, wc.Img, wc.Title, wc.Day,
	)
	if err != nil {
		return err
	}

	return nil
}

func (wc *webComic) read(num int) error {
	db, err := dbConnect()
	if err != nil {
		return err
	}
	defer db.Close()

	row := db.QueryRow(`
		SELECT month, num, link, year, news, safe_title, transcript,
		alt, img, title,day FROM xkcd WHERE num = $1`, num,
	)

	err = row.Scan(&wc.Month, &wc.Num, &wc.Link, &wc.Year, &wc.News, &wc.SafeTitle,
		&wc.Transcript, &wc.Alt, &wc.Img, &wc.Title, &wc.Day,
	)
	if err != nil {
		return err
	}

	return nil
}
