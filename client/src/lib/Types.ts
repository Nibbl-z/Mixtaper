export interface LevelData {
    songName: string,
    songArtist: string,
    description: string,
    uploader: string
}

export interface SearchResult {
    message: LevelData[]
}