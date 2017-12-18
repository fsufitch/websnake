import { ActionReducer } from '@ngrx/store';

import { RootState } from './root.state';
import { loadReducer } from './load';

export const rootReducer: ActionReducer<RootState> = (state=new RootState(), action) => {
  return <RootState>state.merge({
    load: loadReducer(state.load, action),
  });
}
