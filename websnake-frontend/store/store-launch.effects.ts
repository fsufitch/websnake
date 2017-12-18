import { Injectable } from '@angular/core';
import { Actions, Effect } from '@ngrx/effects';

import { StoreInitAction } from './root.actions';

@Injectable()
export class StoreLaunchEffects {
  constructor(private actions$: Actions) {}

  @Effect()
  customStartAction$ = this.actions$
    .ofType('@ngrx/effects/init')
    .map(() => new StoreInitAction());
}
