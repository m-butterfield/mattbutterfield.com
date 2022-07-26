<CsoundSynthesizer>
<CsInstruments>

; the next line sets the volume scale 0-1
; by default this value is 32767
0dbfs = 1

; defines the first instrument
instr 1
; variable for output,  instrument type,  amplitude,  pitch input
;                                                     as parameter 4 in the score
aOut                  vco2              1,          p4
; routes the instrument to default output
out aOut
endin

</CsInstruments>
<CsScore>
; plays three notes in succession
; instrument  time to play at   length to play  frequency to play
i1          0                 1               100
i1          1                 1               200
i1          2                 1               300
</CsScore>
</CsoundSynthesizer>
