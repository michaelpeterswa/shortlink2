import {writable} from "svelte/store";

class Count {
  count: string;

  constructor(data:{count:string}) {
    this.count = data.count;
  }
}
export { Count };
export const count = writable<Count>({"count": "click me!"});