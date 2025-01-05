export interface LevelData {
    songName: string,
    songArtist: string,
    description: string,
    uploader: string
    bpm: number,
    gamesUsed: string[],
    $id: string
}

export interface SearchResult {
    message: LevelData[]
}