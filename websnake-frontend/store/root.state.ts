import { Record } from 'immutable';

import { LoadState } from './load';

export interface RootContainer {
  root: RootState;
}

export interface RootState {
  load: any;
}
export class RootState extends Record({load: undefined}) {}
