export namespace main {
	
	export class Album {
	    id: number;
	    artist: number;
	    name: string;
	    imagePath: string;
	
	    static createFrom(source: any = {}) {
	        return new Album(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.artist = source["artist"];
	        this.name = source["name"];
	        this.imagePath = source["imagePath"];
	    }
	}
	export class Artist {
	    id: number;
	    firstName: string;
	    lastName: string;
	
	    static createFrom(source: any = {}) {
	        return new Artist(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.firstName = source["firstName"];
	        this.lastName = source["lastName"];
	    }
	}
	export class Param {
	    field: string;
	    value: string;
	
	    static createFrom(source: any = {}) {
	        return new Param(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.field = source["field"];
	        this.value = source["value"];
	    }
	}

}

