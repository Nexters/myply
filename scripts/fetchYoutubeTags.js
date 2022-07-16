// # Use global node modules
// npm install -g axios
//  export NODE_PATH=$(npm root --quiet -g)
const fs = require('fs');
const path = require('path');
require('dotenv').config({ path: path.resolve(__dirname, '../.env.local') });
const axios = require('axios');

const API_KEY = process.env.YOUTUBE_API_KEY;

async function main() {
    const musicIds = await getMusicIds();
    const tags = await getVideoTags(musicIds);
    const tagsCsv = arrayToCsvFileContent(tags, 10);

    fs.writeFileSync(path.resolve(__dirname, './youtube_tags.csv'), tagsCsv, { encoding: 'utf-8' });
}

main();

async function getMusicIds() {
    const res = await axios.get('https://www.googleapis.com/youtube/v3/search', {
        params: {
            key: API_KEY,
            part: 'snippet',
            q: 'playlist',
            videoCategoryId: '10',
            type: 'video',
            regionCode: 'kr',
            maxResults: 50
        }
    });

    if (res.status == 200) {
        const tags = res.data.items.map(({ id }) => id.videoId);
        const tagsSet = new Set(tags);

        return [...tagsSet];
    }

    return [];
}

async function getVideoTags(ids) {
    // 'https://www.googleapis.com/youtube/v3/videos?key=AIzaSyCzvd01uELQme8aJWKZ7UgWTKH9ay6vbuk&part=snippet&id='
    const res = await axios.get('https://www.googleapis.com/youtube/v3/videos', {
        params: {
            key: API_KEY,
            part: 'snippet',
            id: ids.join(',')
        }
    });

    if (res.status == 200) {
        return res.data.items.map(item => item.snippet.tags || []).flat();
    }

    return [];
}

function arrayToCsvFileContent(arr, col) {
    let data = '';

    for (let i = 0; i < arr.length; ++i) {
        const d = arr[i];
        
        data += d;
        if (i % (col - 1) === 0) {
            data += '\n';
        } else {
            data += ',';
        }
    }

    return data;
}
