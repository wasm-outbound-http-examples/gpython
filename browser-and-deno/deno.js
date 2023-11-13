import * as _ from './wasm_exec.js';
const go = new window.Go();
const f = await Deno.open('main.wasm');
const buf = await Deno.readAll(f);
const inst = await WebAssembly.instantiate(buf, go.importObject);
go.run(inst.instance);
