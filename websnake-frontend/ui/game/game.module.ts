import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { EffectsModule } from '@ngrx/effects';

import { StoreModule } from 'websnake-frontend/store';

import { GameCanvasComponent } from './game-canvas.component';
import { GameGlobalsService } from './game-globals.service';
import { GameLaunchEffects } from './game-launch.effects';

@NgModule({
  imports: [
    BrowserModule,
    EffectsModule.forFeature([GameLaunchEffects]),
    StoreModule,
  ],
  declarations: [GameCanvasComponent],
  providers: [GameGlobalsService],
  exports: [GameCanvasComponent],
})
export class GameModule {}
