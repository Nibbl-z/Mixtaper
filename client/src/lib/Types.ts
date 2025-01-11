export interface LevelData {
    songName: string,
    songArtist: string,
    description: string,
    uploader: string
    bpm: number,
    gamesUsed: string[],
    $id: string,
    $createdAt: string
}

export interface UserData {
    Username: string,
    DisplayName: string,
    ID: string,
    Bio: string,
    Verified: boolean
}

export interface SearchResult {
    message: LevelData[]
    successful: boolean
}

export interface Game {
    name: string,
    color: string
}


export interface MessageResult {
    message: string
    successful: boolean
}

export interface LoginResult {
    message: string
    successful: boolean
    secret?: string
}

export interface GetLevelResult {
    message: LevelData
    successful: boolean
}

export interface PostResult {
    message: string
    successful: boolean
    id?: string
}

export interface GetUserResult {
    message: UserData,
    successful: boolean
}