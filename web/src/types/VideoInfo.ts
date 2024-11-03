export interface VideoInfo {
  desc: string
  author: Author
  statsV2: StatsV2
  relatedVideos: RelatedVideo[]
}

export interface RelatedVideo {
  url: string
  title: string
  username: string
  thumbnail: string
}

export interface Author {
  id: string
  shortId: string
  uniqueId: string
  nickname: string
  avatarLarger: string
  avatarMedium: string
  avatarThumb: string
  signature: string
}

export interface StatsV2 {
  diggCount: string
  shareCount: string
  commentCount: string
  playCount: string
  collectCount: string
  repostCount: string
}
