import { Injectable } from '@angular/core';
import { Effect, Actions } from '@ngrx/effects';
import { Observable } from 'rxjs';

import { StoreInitAction } from 'websnake-frontend/store';
import { GameGlobalsService } from './game-globals.service';

@Injectable()
export class GameLaunchEffects {
  constructor(
    private actions$: Actions,
    private gameGlobalsService: GameGlobalsService,
  ) {}

  @Effect()
  startPixiApplication$ = Observable.of(null).do(() => {
      this.gameGlobalsService.initializePixiApplication();
      this.gameGlobalsService.initializeResources();
    })
    .flatMap(() => Observable.empty());
}
