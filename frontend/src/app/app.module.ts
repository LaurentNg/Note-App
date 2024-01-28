import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { DividerModule } from 'primeng/divider';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { InputTextareaModule } from 'primeng/inputtextarea';
import { PanelModule } from 'primeng/panel';
import { TabViewModule } from 'primeng/tabview';
import { TooltipModule } from 'primeng/tooltip';

import {ClipboardModule} from '@angular/cdk/clipboard';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './pages/app/app.component';
import { NotesComponent } from './pages/notes/notes.component';
import { EncryptTextComponent } from './components/encrypt-text/encrypt-text.component';
import { DecryptTextComponent } from './components/decrypt-text/decrypt-text.component';

@NgModule({
  declarations: [
    AppComponent,
    NotesComponent,
    EncryptTextComponent,
    DecryptTextComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    ClipboardModule,
    ButtonModule,
    CardModule,
    DividerModule,
    FormsModule,
    InputTextareaModule,
    ReactiveFormsModule,
    PanelModule,
    TabViewModule,
    TooltipModule,
    ClipboardModule,
    AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
