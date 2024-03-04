class Post{
    content :string;
    privacy :string;
    followersSelectedId? : string[];
    creationDate? : Date;
    imageUrl? : string;
    userId? : string;
    constructor(data : object) {
        this.content  = data['content'];
        this.privacy = data['privacy'];
        this.creationDate = data['creationDate'];
        this.imageUrl = data['imageUrl'];
        this.userId = data['userId'];
    }
    validate(): [boolean, string] {

        if (this.content.length > 800 ) {
            return [false, "Post content should not exceed 800 characters"];
        } 
        //TODO: add  more case handling 
    
        return [true, "Post data is valid"];
    }
}