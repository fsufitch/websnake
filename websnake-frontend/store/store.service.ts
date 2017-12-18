import { Injectable } from '@angular/core';
import { Store, Action } from '@ngrx/store';

import { RootContainer } from './root.state';
import { LoadSelector } from './load';

@Injectable()
export class StoreService {
  constructor(private store: Store<RootContainer>) {}

  dispatch(action: Action) {
    this.store.dispatch(action);
  }

  getRootState() {
    return this.store.let(store$ => store$.select(s => s.root))
      .filter(state => state !== undefined);
  }

  selectLoad() {
    return new LoadSelector(this.getRootState().select(s => s.load));
  }
}
