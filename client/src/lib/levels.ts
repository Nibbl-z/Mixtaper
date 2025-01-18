import { PUBLIC_PROJECT_ID } from "$env/static/public";
import type { Game, LevelData, SearchResult } from "./Types";

async function checkCoverArt(url: string): Promise<boolean> {
    const response = await fetch(url)
    
    return response.ok
}

export async function getCoverArtById(id: string): Promise<string> {
    const coverUrl = `https://cloud.appwrite.io/v1/storage/buckets/cover_art/files/${id}/view?project=${PUBLIC_PROJECT_ID}&project=${PUBLIC_PROJECT_ID}`
    
    return checkCoverArt(coverUrl).then(exists => (
        exists ? coverUrl : "/placeholder_coverart.png"
    ))
}

export async function loadLevels(levels: SearchResult): Promise<{ results: LevelData[], covers: Record<string, string> }> {
    let results: LevelData[] = [];
    let covers: Record<string, string> = {}
    
    results = levels.message
    
    const coverPromises = results.map(result => {
        const coverUrl = `https://cloud.appwrite.io/v1/storage/buckets/cover_art/files/${result.$id}/view?project=${PUBLIC_PROJECT_ID}&project=${PUBLIC_PROJECT_ID}`
        return checkCoverArt(coverUrl).then(exists => ({
            id: result.$id,
            url: exists ? coverUrl : "/placeholder_coverart.png"
        }))
    })
    
    const loadedCovers = await Promise.all(coverPromises)
    covers = loadedCovers.reduce((acc: Record<string, string>, cover) => {
        acc[cover.id] = cover.url
        return acc
    }, {})

    return { results, covers }
}

export const gameNames: Record<string, Game> = {
    "flipperSnapper" : {
        name : "Flipper Snapper",
        color : "#D68D62"
    },
    "meetAndTweet" : {
        name : "Meet & Tweet",
        color : "#19a1e0"
    },
    "hammerTime" : {
        name : "Hammer Time",
        color : "#736c64"
    }
}