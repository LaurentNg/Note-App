import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-notes',
  templateUrl: './notes.component.html',
  styleUrls: ['./notes.component.scss']
})
export class NotesComponent implements OnInit {
  notesToEncrypt = new FormGroup({
    text: new FormControl<string>('', [Validators.required]),
    passphrase: new FormControl<string>('', [Validators.required])
  });
  constructor() { }

  ngOnInit(): void {
  }

  onSubmit(): void {
    if (this.notesToEncrypt.valid) {
      const text = this.notesToEncrypt.get('text')?.value
      const passphrase = this.notesToEncrypt.get('passphrase')?.value
      console.log(text)
      console.log(passphrase)
    }
    else console.log('not valid')
  }

}
