# server-myself

## 게시물 JSON
{
    profile: {
        profileImg: string,
        level: Int,
        nickname: string,
    },
    title: string,
    uploadTime: string,
    contents : [
        contents1: {
            imgs: [string],
            text: string,
        }
    ],
    comments: [
        comment1: {
            profile: {
                profileImg: string,
                level: Int,
                nickname: string,
            },
            text: string,
            uploadTime: string,
        }
    ],
    
}