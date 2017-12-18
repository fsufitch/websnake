import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app/app.component';
import { GameModule } from './game';

@NgModule({
  imports: [BrowserModule, GameModule],
  declarations: [AppComponent],
  bootstrap: [AppComponent],
})
export class UIModule {}
