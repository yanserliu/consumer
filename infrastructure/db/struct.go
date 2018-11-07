package db

type User struct {
	Id          int64
	Email       string
	Password    string
	Nickname    string
	Realname    string
	Avatar      string
	Avatar_min  string
	Avatar_max  string
	Province    string
	City        string
	Company     string
	Address     string
	Postcode    string
	Mobile      string
	Website     string
	Sex         int64
	Qq          string
	Msn         string
	Weibo       string
	Ctype       int64
	Role        int64
	Hotness     float64
	Hotup       int64
	Hotdown     int64
	Hotscore    int64
	Views       int64
	LastLoginIp string
	LoginCount  int64
}

//category,Pid:root
type Category struct {
	Id         int64
	Pid        int64
	Uid        int64
	Ctype      int64
	Title      string
	Content    string
	Attachment string

	Hotness  float64
	Hotup    int64
	Hotdown  int64
	Hotscore int64
	Views    int64
	Author   string

	NodeCount      int64
	NodeLastUserId int64
}

//node,Pid:category
type Node struct {
	Id         int64
	Pid        int64
	Uid        int64
	Ctype      int64
	Title      string
	Content    string
	Attachment string

	Hotness  float64
	Hotup    int64
	Hotdown  int64
	Hotscore int64
	Views    int64
	Author   string

	TopicCount      int64
	TopicLastUserId int64
}

//topic,Pid:node
type Topic struct {
	Id         int64
	Cid        int64
	Nid        int64
	Uid        int64
	Ctype      int64
	Title      string
	Content    string
	Attachment string

	Hotness         float64
	Hotup           int64
	Hotdown         int64
	Hotscore        int64
	Views           int64
	Author          string
	ReplyCount      int64
	ReplyLastUserId int64
}

//reply,Pid:topic
type Reply struct {
	Id         int64
	Uid        int64
	Pid        int64 //Topic id
	Ctype      int64
	Content    string
	Attachment string

	Hotness  float64
	Hotup    int64
	Hotdown  int64
	Hotscore int64
	Views    int64
	Author   string
	Email    string
	Website  string
}

type File struct {
	Id       int64
	Cid      int64
	Nid      int64
	Uid      int64
	Pid      int64
	Ctype    int64
	Filename string
	Content  string
	Hash     string
	Location string
	Url      string
	Size     int64

	Hotness  float64
	Hotup    int64
	Hotdown  int64
	Hotscore int64
	Views    int64

	ReplyCount      int64
	ReplyLastUserId int64
}
