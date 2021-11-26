import {writable} from "svelte/store";

class Shortlink {
  shortid: string;
  shortlink: string;
  url: string;

  constructor(data:{shortid:string, shortlink:string, url:string}) {
    this.shortid = data.shortid;
    this.shortlink = data.shortlink;
    this.url = data.url;
  }
}
export { Shortlink };
export const shortlink = writable<Shortlink>(null);