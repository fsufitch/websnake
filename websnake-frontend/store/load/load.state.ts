import { Record } from 'immutable';

export interface LoadState {
  canvasReady: boolean;
  resourcesReady: boolean;
}

export class LoadState extends Record({
  canvasReady: false,
  resourcesReady: false,
}) {}
