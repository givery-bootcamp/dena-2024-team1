import { atom } from "jotai";

import { Sketch } from "~/generated";

export const selectedSketchAtom = atom<Sketch | null>(null);
