export namespace api {
	
	export class TaskServiceWrapper {
	    Service?: tasks.Service;
	
	    static createFrom(source: any = {}) {
	        return new TaskServiceWrapper(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Service = this.convertValues(source["Service"], tasks.Service);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace tasks {
	
	export class TasksService {
	
	
	    static createFrom(source: any = {}) {
	        return new TasksService(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class TasklistsService {
	
	
	    static createFrom(source: any = {}) {
	        return new TasklistsService(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class Service {
	    BasePath: string;
	    UserAgent: string;
	    // Go type: TasklistsService
	    Tasklists?: any;
	    // Go type: TasksService
	    Tasks?: any;
	
	    static createFrom(source: any = {}) {
	        return new Service(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.BasePath = source["BasePath"];
	        this.UserAgent = source["UserAgent"];
	        this.Tasklists = this.convertValues(source["Tasklists"], null);
	        this.Tasks = this.convertValues(source["Tasks"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

