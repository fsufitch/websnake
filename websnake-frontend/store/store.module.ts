import { NgModule } from '@angular/core';
import { StoreModule as NgrxStoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';

import { RootState } from './root.state';
import { StoreService } from './store.service';
import { StoreLaunchEffects } from './store-launch.effects';
import { rootReducer } from './root.reducer';

@NgModule({
  imports: [
    NgrxStoreModule.forRoot({root: rootReducer}),
    EffectsModule.forRoot([StoreLaunchEffects]),
  ],
  providers: [StoreService],
})
export class StoreModule {}
