import { ActionReducer } from '@ngrx/store';

import * as loadActions from './load.actions';
import { LoadState } from './load.state';

export const loadReducer: ActionReducer<LoadState> = (state=new LoadState(), action) => {
  switch(action.type) {
    case loadActions.CanvasReadyAction.type: {
      state = <LoadState>state.set('canvasReady', true);
      break;
    }
    case loadActions.ResourcesReadyAction.type: {
      state = <LoadState>state.set('resourcesReady', true);
      break;
    }
  }
  return state;
};
