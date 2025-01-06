import type { LevelData, SearchResult } from "./Types";

async function checkCoverArt(url: string): Promise<boolean> {
    const response = await fetch(url)
    
    return response.ok
}

export async function loadLevels(levels: SearchResult): Promise<{ results: LevelData[], covers: Record<string, string> }> {
    let results: LevelData[] = [];
    let covers: Record<string, string> = {}
    
    results = levels.message
    
    const coverPromises = results.map(result => {
        const coverUrl = `https://cloud.appwrite.io/v1/storage/buckets/cover_art/files/${result.$id}/view?project=676f205d000370a15786&project=676f205d000370a15786`
        return checkCoverArt(coverUrl).then(exists => ({
            id: result.$id,
            url: exists ? coverUrl : "/PLACEHOLDER.png"
        }))
    })
    
    const loadedCovers = await Promise.all(coverPromises)
    covers = loadedCovers.reduce((acc: Record<string, string>, cover) => {
        acc[cover.id] = cover.url
        return acc
    }, {})

    return { results, covers }
}